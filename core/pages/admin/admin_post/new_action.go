package admin_post

import (
	"github.com/daqing/waterway/core/api/node_api"
	"github.com/daqing/waterway/lib/page_resp"
	"github.com/daqing/waterway/lib/repo"
	"github.com/gin-gonic/gin"
)

func NewAction(c *gin.Context) {
	nodes, err := repo.Find[node_api.Node](
		[]string{"id", "name"},
		[]repo.KVPair{},
	)

	if err != nil {
		page_resp.Error(c, err)
		return
	}

	data := map[string]any{
		"Nodes": nodes,
	}

	page_resp.Page(c, "core", "admin.post", "new", data)
}
