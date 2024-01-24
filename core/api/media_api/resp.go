package media_api

import "github.com/daqing/waterway/lib/repo"

type MediaResp struct {
	Id int64

	CreatedAt repo.Timestamp
	UpdatedAt repo.Timestamp
}

func (r MediaResp) Fields() []string {
	return []string{"id"}
}
