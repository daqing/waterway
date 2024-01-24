package checkin_api

import (
	"time"

	"github.com/daqing/waterway/lib/repo"
)

type Checkin struct {
	repo.Model

	UserId repo.IdType

	Year  int
	Month time.Month
	Day   int

	Acc int // 连续签到次数
}

const tableName = "checkin"

func (c Checkin) TableName() string { return tableName }
