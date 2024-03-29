package admin_menu

import (
	"github.com/daqing/waterway/core/pages/admin/helper"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	g := r.Group("/menu")
	{
		g.GET("", helper.CheckAdmin(IndexAction))
		g.GET("/new", helper.CheckAdmin(NewAction))
		g.POST("/create", helper.CheckAdmin(CreateAction))

		g.GET("/edit", helper.CheckAdmin(EditAction))
		g.POST("/update", helper.CheckAdmin(UpdateAction))

		// TODO: csrf protection
		g.GET("/delete", helper.CheckAdmin(DeleteAction))
	}
}
