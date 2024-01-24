package action_api

import "github.com/daqing/waterway/lib/repo"

type Action struct {
	repo.Model

	UserId     repo.IdType
	Action     string
	TargetType string
	TargetId   repo.IdType
}

const ActionLike = "like"
const ActionFollow = "follow"
const ActionFavorite = "favorite"

const actionTableName = "actions"

func (a Action) TableName() string { return actionTableName }
