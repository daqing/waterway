package admin_node

import (
	"fmt"

	"github.com/daqing/waterway/core/api/node_api"
	"github.com/daqing/waterway/lib/page_resp"
	"github.com/daqing/waterway/lib/utils"
	"github.com/gin-gonic/gin"
)

type CreateParams struct {
	Name  string `form:"name"`
	Key   string `form:"key"`
	Place string `form:"place"`
}

func CreateAction(c *gin.Context) {
	var p CreateParams

	if err := c.ShouldBind(&p); err != nil {
		page_resp.Error(c, err)
		return
	}

	name := utils.TrimFull(p.Name)
	key := utils.TrimFull(p.Key)

	if len(name) == 0 || len(key) == 0 {
		page_resp.Error(c, fmt.Errorf("name or key must not be empty"))
		return
	}

	_, err := node_api.CreateNode(name, key, p.Place, 0, 0)
	if err != nil {
		page_resp.Error(c, err)
		return
	}

	page_resp.Redirect(c, "/admin/node")
}
