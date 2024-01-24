package media_api

import (
	"time"

	"github.com/daqing/waterway/lib/repo"
)

type MediaFile struct {
	repo.Model

	UserId    repo.IdType
	Filename  string
	Mime      string
	Bytes     repo.IdType
	ExpiredAt time.Time
}

const tableName = "media_files"

func (m MediaFile) TableName() string { return tableName }
