package comment_api

import (
	"github.com/daqing/waterway/core/api/user_api"
	"github.com/daqing/waterway/lib/repo"
)

func CreateComment(currentUser *user_api.User, polyModel repo.PolyModel, content string) (*Comment, error) {
	return createComment(currentUser, polyModel.PolyType(), polyModel.PolyId(), content)
}

func createComment(currentUser *user_api.User, targetType string, targetId repo.IdType, content string) (*Comment, error) {
	return repo.Insert[Comment]([]repo.KVPair{
		repo.KV("user_id", currentUser.ID),
		repo.KV("target_type", targetType),
		repo.KV("target_id", targetId),
		repo.KV("content", content),
	})
}
