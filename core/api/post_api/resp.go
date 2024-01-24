package post_api

import "github.com/daqing/waterway/lib/repo"

type PostResp struct {
	Id         repo.IdType
	UserId     repo.IdType
	NodeId     repo.IdType
	Title      string
	CustomPath string
	Place      string
	Content    string
	Fee        int
	CreatedAt  repo.Timestamp
	UpdatedAt  repo.Timestamp
}

func (pr PostResp) Fields() []string {
	return []string{
		"id", "user_id", "node_id",
		"title", "custom_path",
		"place", "content", "fee",
	}
}
