package setting_api

import "github.com/daqing/waterway/lib/repo"

type SettingResp struct {
	Id int64

	Key string
	Val string

	CreatedAt repo.Timestamp
	UpdatedAt repo.Timestamp
}

func (sr SettingResp) Fields() []string {
	return []string{"id", "key", "val"}
}
