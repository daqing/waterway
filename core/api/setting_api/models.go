package setting_api

import "github.com/daqing/waterway/lib/repo"

type Setting struct {
	repo.Model

	Key string
	Val string
}

const tableName = "settings"

func (s Setting) TableName() string { return tableName }
