package helper

import (
	"github.com/daqing/waterway/core/api/user_api"
	"github.com/daqing/waterway/lib/page_resp"
	"github.com/daqing/waterway/lib/utils"
	"github.com/gin-gonic/gin"
)

func CheckAdmin(action gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.CookieToken(c)
		if err != nil {
			page_resp.Redirect(c, "/")
			return
		}

		var admin = user_api.CurrentAdmin(token)
		if admin == nil {
			page_resp.Redirect(c, "/session/sign_in")
			return
		}

		c.Set("admin", admin)

		action(c)
	}
}
