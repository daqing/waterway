package user_api

import (
	"github.com/daqing/waterway/core/api/membership_api"
	"github.com/daqing/waterway/lib/repo"
)

type User struct {
	repo.Model

	Nickname          string
	Username          string
	Phone             string
	Email             string
	Avatar            string
	Role              UserRole
	APIToken          string
	EncryptedPassword string
	Balance           repo.PriceCent
}

const tableName = "users"

func (u User) TableName() string { return tableName }

type UserRole int

const (
	AllRole UserRole = iota
	RootRole
	AdminRole
	BasicRole
)

const polyType = "user"

func (u *User) PolyType() string { return polyType }

func (u *User) PolyId() repo.IdType { return u.ID }

func (u *User) Membership() (*membership_api.MembershipResp, error) {
	return membership_api.MembershipFor(u.ID)
}

func (u *User) IsAdmin() bool { return u.Role == AdminRole || u.Role == RootRole }

func RoleName(role UserRole) string {
	switch role {
	case RootRole:
		return "ROOT"
	case AdminRole:
		return "ADMIN"
	case BasicRole:
		return "BASIC"
	default:
		return "[OTHER]"
	}
}
