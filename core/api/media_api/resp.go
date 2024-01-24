package media_api

import "github.com/daqing/waterway/lib/repo"

type MediaResp struct {
	Id repo.IdType

	CreatedAt repo.Timestamp
	UpdatedAt repo.Timestamp
}

func (r MediaResp) Fields() []string {
	return []string{"id"}
}
