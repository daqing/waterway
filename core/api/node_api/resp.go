package node_api

import (
	"github.com/daqing/waterway/lib/repo"
)

type NodeResp struct {
	Id repo.IdType

	Name      string
	Key       string
	ParentKey string
	Level     int

	CreatedAt repo.Timestamp
	UpdatedAt repo.Timestamp
}

func (ur NodeResp) Fields() []string {
	return []string{"id", "name", "key", "parent_id", "level"}
}
