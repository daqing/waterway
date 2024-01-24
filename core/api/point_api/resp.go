package point_api

import "github.com/daqing/waterway/lib/repo"

type PointResp struct {
	Id repo.IdType

	UserId repo.IdType
	Count  int

	CreatedAt repo.Timestamp
	UpdatedAt repo.Timestamp
}

func (r PointResp) Fields() []string {
	return []string{"id", "user_id", "count"}
}
