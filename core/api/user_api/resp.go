package user_api

import "github.com/daqing/waterway/lib/repo"

type UserResp struct {
	Id        repo.IdType
	Nickname  string
	Username  string
	ApiToken  string
	Role      UserRole
	CreatedAt repo.Timestamp
	UpdatedAt repo.Timestamp
}

func (ur UserResp) Fields() []string {
	return []string{"id", "username", "nickname", "role", "api_token"}
}
