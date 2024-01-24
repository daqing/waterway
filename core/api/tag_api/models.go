package tag_api

import "github.com/daqing/waterway/lib/repo"

type Tag struct {
	repo.Model

	Name string
}

const tableName = "tags"

func (t Tag) TableName() string { return tableName }

type TagRelation struct {
	TagId        repo.IdType
	RelationType string
	RelationId   repo.IdType
}

const relationTableName = "tag_relation"

func (tr TagRelation) TableName() string { return relationTableName }
