package admin_menu

import (
	"github.com/daqing/waterway/core/api/menu_api"
	"github.com/daqing/waterway/lib/page_resp"
	"github.com/daqing/waterway/lib/utils"
	"github.com/gin-gonic/gin"
)

func EditAction(c *gin.Context) {
	id := utils.TrimFull(c.Query("id"))

	if len(id) == 0 {
		page_resp.Redirect(c, "/admin/menu")
	}

	menu, err := menu_api.FindBy("id", id)
	if err != nil {
		page_resp.Error(c, err)
		return
	}

	data := map[string]any{
		"Menu": menu,
	}

	page_resp.Page(c, "core", "admin.menu", "edit", data)
}
