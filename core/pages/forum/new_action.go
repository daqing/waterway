package forum

import (
	"time"

	"github.com/daqing/waterway/core/api/node_api"
	"github.com/daqing/waterway/core/api/user_api"
	"github.com/daqing/waterway/lib/page_resp"
	"github.com/daqing/waterway/lib/repo"
	"github.com/daqing/waterway/lib/utils"
	"github.com/gin-gonic/gin"
)

func NewAction(c *gin.Context) {
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
				Id:   node.Id,
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
		"Session":  SessionData(currentUser),
	}

	page_resp.Page(c, "core", "forum!", "new", data)
}
