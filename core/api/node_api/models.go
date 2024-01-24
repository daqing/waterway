package node_api

import "github.com/daqing/waterway/lib/repo"

type Node struct {
	repo.Model

	ParentId repo.IdType

	Name string
	Key  string

	Level int
	Place string
}

const tableName = "nodes"

func (n Node) TableName() string { return tableName }
