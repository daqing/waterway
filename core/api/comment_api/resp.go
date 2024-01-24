package comment_api

import "github.com/daqing/waterway/lib/repo"

type CommentResp struct {
	Id repo.IdType

	TargetId   repo.IdType
	TargetType string
	Content    string

	CreatedAt repo.Timestamp
	UpdatedAt repo.Timestamp
}

func (ur CommentResp) Fields() []string {
	return []string{"id", "target_id", "target_type", "content"}
}
