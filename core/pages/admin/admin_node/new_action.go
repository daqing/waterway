package admin_node

import (
	"github.com/daqing/waterway/lib/page_resp"
	"github.com/gin-gonic/gin"
)

func NewAction(c *gin.Context) {
	page_resp.Page(c, "core", "admin.node", "new", nil)
}
