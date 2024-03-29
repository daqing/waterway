package admin_post

import (
	"github.com/daqing/waterway/core/api/node_api"
	"github.com/daqing/waterway/core/api/post_api"
	"github.com/daqing/waterway/lib/page_resp"
	"github.com/daqing/waterway/lib/repo"
	"github.com/gin-gonic/gin"
)

func EditAction(c *gin.Context) {
	id := c.Query("id")

	post, err := repo.FindRow[post_api.Post](
		[]string{"id", "title", "custom_path", "place", "content"},
		[]repo.KVPair{
			repo.KV("id", id),
		},
	)

	if err != nil {
		page_resp.Error(c, err)
		return
	}

	nodes, err := repo.Find[node_api.Node](
		[]string{"id", "name"},
		[]repo.KVPair{},
	)

	if err != nil {
		page_resp.Error(c, err)
		return
	}

	data := map[string]any{
		"Post":  post,
		"Nodes": nodes,
	}

	page_resp.Page(c, "core", "admin.post", "edit", data)
}
