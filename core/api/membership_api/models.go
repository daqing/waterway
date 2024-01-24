package membership_api

import (
	"time"

	"github.com/daqing/waterway/lib/repo"
)

type Membership struct {
	repo.Model

	UserId    repo.IdType
	Name      string
	ExpiredAt time.Time
}

const tableName = "memberships"

func (m Membership) TableName() string { return tableName }

type MembershipType string

const Writer MembershipType = "writer"
