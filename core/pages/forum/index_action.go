package forum

import (
	"fmt"
	"time"

	"github.com/daqing/waterway/core/api/media_api"
	"github.com/daqing/waterway/core/api/node_api"
	"github.com/daqing/waterway/core/api/post_api"
	"github.com/daqing/waterway/core/api/user_api"
	"github.com/daqing/waterway/lib/page_resp"
	"github.com/daqing/waterway/lib/repo"
	"github.com/daqing/waterway/lib/utils"
	"github.com/gin-gonic/gin"
)

type PostItem struct {
	Id        int64
	Title     string
	Url       string
	TimeAgo   string
	UserName  string
	NodeName  string
	AvatarURL string
}

type NodeItem struct {
	Id   int64
	Name string
	URL  string
}

func IndexAction(c *gin.Context) {
	posts, err := post_api.Posts(
		[]string{"id", "title", "custom_path", "user_id", "node_id"},
		"forum", // TODO: define constant
		"id DESC",
		0,
		50,
	)

	if err != nil {
		page_resp.Error(c, err)
		return
	}

	nameMap, err := node_api.NameMap("forum")

	if err != nil {
		page_resp.Error(c, err)
		return
	}

	postsShow := []*PostItem{}

	for _, post := range posts {
		url := fmt.Sprintf("/forum/post/%d", post.Id)

		if len(post.CustomPath) > 0 {
			url = fmt.Sprintf("/forum/post/%s", post.CustomPath)
		}

		postsShow = append(postsShow,
			&PostItem{
				Id:        post.Id,
				Title:     post.Title,
				Url:       url,
				TimeAgo:   utils.TimeAgo(post.CreatedAt),
				NodeName:  nameMap[post.NodeId],
				UserName:  user_api.Nickname(post.UserId),
				AvatarURL: media_api.AssetHostPath(post.UserAvatar()),
			},
		)
	}

	nodes, err := repo.Find[node_api.Node](
		[]string{"id", "name", "key"},
		[]repo.KVPair{
			repo.KV("place", "forum"),
		},
	)

	if err != nil {
		page_resp.Error(c, err)
		return
	}

	rootPath := utils.PathPrefix("forum")

	nodeItems := []*NodeItem{}

	for _, node := range nodes {
		nodeItems = append(nodeItems,
			&NodeItem{
				Name: node.Name,
				URL:  "/forum/node/" + node.Key,
			})
	}

	token, _ := utils.CookieToken(c)
	currentUser := user_api.CurrentUser(token)

	data := map[string]any{
		"Title":    ForumTitle(),
		"Tagline":  ForumTagline(),
		"Year":     time.Now().Year(),
		"RootPath": rootPath,
		"Nodes":    nodeItems,
		"Posts":    postsShow,
		"Session":  SessionData(currentUser),
	}

	page_resp.Page(c, "core", "forum!", "index", data)
}
