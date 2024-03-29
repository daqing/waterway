package comment_api

import (
	"github.com/daqing/waterway/core/api/user_api"
	"github.com/daqing/waterway/lib/repo"
)

type Comment struct {
	repo.Model

	UserId repo.IdType

	TargetId   repo.IdType
	TargetType string
	Content    string
}

const tableName = "comments"

func (c Comment) TableName() string { return tableName }

func (c *Comment) User() *user_api.User {
	user, err := repo.FindRow[user_api.User](
		[]string{"id", "nickname", "username", "avatar"},
		[]repo.KVPair{
			repo.KV("id", c.UserId),
		},
	)

	if err != nil {
		return nil
	}

	return user
}
