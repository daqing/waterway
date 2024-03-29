package admin_node

import (
	"github.com/daqing/waterway/core/api/node_api"
	"github.com/daqing/waterway/lib/page_resp"
	"github.com/daqing/waterway/lib/repo"
	"github.com/gin-gonic/gin"
)

func EditAction(c *gin.Context) {
	id := c.Query("id")

	node, err := repo.FindRow[node_api.Node](
		[]string{"id", "name", "key"},
		[]repo.KVPair{
			repo.KV("id", id),
		},
	)

	if err != nil {
		page_resp.Error(c, err)
		return
	}

	data := map[string]any{
		"Node": node,
	}

	page_resp.Page(c, "core", "admin.node", "edit", data)
}
