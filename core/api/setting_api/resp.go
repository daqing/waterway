package setting_api

import "github.com/daqing/waterway/lib/repo"

type SettingResp struct {
	Id repo.IdType

	Key string
	Val string

	CreatedAt repo.Timestamp
	UpdatedAt repo.Timestamp
}

func (sr SettingResp) Fields() []string {
	return []string{"id", "key", "val"}
}
