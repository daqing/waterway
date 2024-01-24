package checkin_api

import (
	"time"

	"github.com/daqing/waterway/lib/repo"
)

type CheckinResp struct {
	Id repo.IdType

	UserId repo.IdType

	Year  int
	Month time.Month
	Day   int

	Acc int // 连续签到次数

	CreatedAt repo.Timestamp
	UpdatedAt repo.Timestamp
}

func (c CheckinResp) Fields() []string {
	return []string{"id", "user_id", "year", "month", "day", "acc"}
}
