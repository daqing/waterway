package post_api

import (
	"github.com/daqing/waterway/lib/api_resp"
	"github.com/daqing/waterway/lib/repo"
	"github.com/gin-gonic/gin"
)

type ShowParams struct {
	ID repo.IdType `form:"id"`
}

func ShowAction(c *gin.Context) {
	var p ShowParams

	if err := c.ShouldBind(&p); err != nil {
		api_resp.LogError(c, err)
		return
	}

	post, err := repo.FindRow[Post]([]string{
		"id", "user_id", "node_id", "title", "content",
	}, []repo.KVPair{
		repo.KV("id", p.ID),
	})

	if err != nil {
		api_resp.LogError(c, err)
		return
	}

	api_resp.OK(c, gin.H{"post": post})
}
