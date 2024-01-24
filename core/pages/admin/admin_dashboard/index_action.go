package admin_dashboard

import (
	"github.com/daqing/waterway/lib/page_resp"
	"github.com/gin-gonic/gin"
)

func IndexAction(c *gin.Context) {
	data := map[string]any{}
	page_resp.Page(c, "core", "admin.dashboard", "index", data)
}
