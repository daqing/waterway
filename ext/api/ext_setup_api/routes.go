package ext_setup_api

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	g := r.Group("/setup")
	{
		g.POST("/create", CreateAction)
	}
}
