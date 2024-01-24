package forum

import (
	"fmt"

	"github.com/daqing/waterway/core/api/comment_api"
	"github.com/daqing/waterway/core/api/post_api"
	"github.com/daqing/waterway/core/api/user_api"
	"github.com/daqing/waterway/lib/page_resp"
	"github.com/daqing/waterway/lib/repo"
	"github.com/daqing/waterway/lib/utils"
	"github.com/gin-gonic/gin"
)

type CreateCommentParams struct {
	TargetId repo.IdType `form:"target_id"`
	Content  string      `form:"content"`
}

func CreateCommentAction(c *gin.Context) {
	var p CreateCommentParams

	if err := c.ShouldBind(&p); err != nil {
		page_resp.Error(c, err)
		return
	}

	token, _ := utils.CookieToken(c)
	currentUser := user_api.CurrentUser(token)
	if currentUser == nil {
		page_resp.Redirect(c, "/forum")
		return
	}

	polyModel := &post_api.Post{}
	polyModel.ID = p.TargetId

	_, err := comment_api.CreateComment(currentUser, polyModel, p.Content)
	if err != nil {
		page_resp.Error(c, err)
		return
	}

	backPath := fmt.Sprintf("/forum/post/%d", p.TargetId)
	page_resp.Redirect(c, backPath)
}
