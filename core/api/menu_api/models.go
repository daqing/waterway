package menu_api

import "github.com/daqing/waterway/lib/repo"

type Menu struct {
	repo.Model

	Name  string
	URL   string
	Place string
}

const tableName = "menus"

func (m Menu) TableName() string { return tableName }
