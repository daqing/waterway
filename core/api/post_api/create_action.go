package post_api

import (
	"strings"

	"github.com/daqing/waterway/core/api/user_api"
	"github.com/daqing/waterway/lib/api_resp"
	"github.com/daqing/waterway/lib/repo"
	"github.com/gin-gonic/gin"
)

type CreateParams struct {
	NodeId  repo.IdType `json:"node_id"`
	Title   string      `json:"title"`
	Place   string      `json:"place"`
	Content string      `json:"content"`
	Fee     int         `json:"fee"`
	Tags    string      `json:"tags"` // 使用英文逗号分隔
}

func CreateAction(c *gin.Context) {
	var p CreateParams

	if err := c.BindJSON(&p); err != nil {
		api_resp.LogError(c, err)
		return
	}

	user := user_api.CurrentUser(c.GetHeader("X-Auth-Token"))

	if user == nil {
		api_resp.LogInvalidUser(c)
		return
	}

	tags := strings.Split(p.Tags, ",")

	// TODO: add custom path parameters
	post, err := CreatePost(p.Title, "", p.Place, p.Content, user.ID, p.NodeId, p.Fee, tags)
	if err != nil {
		api_resp.LogError(c, err)
		return
	}

	item := repo.ItemResp[Post, PostResp](post)

	api_resp.OK(c, gin.H{"post": item})
}
