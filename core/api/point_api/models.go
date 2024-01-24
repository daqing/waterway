package point_api

import "github.com/daqing/waterway/lib/repo"

// 用户积分
type Point struct {
	repo.Model

	UserId repo.IdType
	Count  int
}

const tableName = "points"

func (m Point) TableName() string { return tableName }
