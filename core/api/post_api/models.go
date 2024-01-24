package post_api

import (
	"github.com/daqing/waterway/core/api/comment_api"
	"github.com/daqing/waterway/core/api/node_api"
	"github.com/daqing/waterway/core/api/user_api"
	"github.com/daqing/waterway/lib/repo"
)

type Post struct {
	repo.Model

	UserId     repo.IdType
	NodeId     repo.IdType
	Title      string
	CustomPath string
	Place      string
	Content    string
	Fee        int
}

const tableName = "posts"

func (p Post) TableName() string { return tableName }

const polyType = "post"

func (p *Post) PolyId() repo.IdType { return p.ID }
func (p *Post) PolyType() string    { return polyType }

func (p *Post) User() *user_api.User {
	user, err := repo.FindRow[user_api.User](
		[]string{"id", "nickname", "username", "avatar"},
		[]repo.KVPair{
			repo.KV("id", p.UserId),
		},
	)

	if err != nil {
		return nil
	}

	return user
}

func (p *Post) UserAvatar() string {
	user := p.User()

	if user == nil {
		return repo.EMPTY_STRING
	}

	return user.Avatar
}

func (p *Post) Node() *node_api.Node {
	node, err := repo.FindRow[node_api.Node](
		[]string{"id", "name", "key"},
		[]repo.KVPair{
			repo.KV("id", p.NodeId),
		},
	)

	if err != nil {
		return nil
	}

	return node
}

func (p *Post) Comments() ([]*comment_api.Comment, error) {
	return repo.Find[comment_api.Comment](
		[]string{"id", "user_id", "content"},
		[]repo.KVPair{
			repo.KV("target_type", p.PolyType()),
			repo.KV("target_id", p.PolyId()),
		},
	)
}
