package admin_post

import (
	"github.com/daqing/waterway/core/api/post_api"
	"github.com/daqing/waterway/lib/page_resp"

	"github.com/gin-gonic/gin"
)

type IndexParams struct {
	Page int `form:"page"`
}

func IndexAction(c *gin.Context) {
	var p IndexParams

	if err := c.Bind(&p); err != nil {
		page_resp.Error(c, err)
		return
	}

	posts, err := post_api.Posts(
		[]string{"id", "title", "custom_path", "place", "user_id", "node_id"},
		"",
		"id DESC",
		p.Page,
		50,
	)

	if err != nil {
		page_resp.Error(c, err)
		return
	}

	data := map[string]any{
		"Posts": posts,
	}

	page_resp.Page(c, "core", "admin.post", "index", data)
}
