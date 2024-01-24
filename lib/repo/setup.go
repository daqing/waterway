package repo

import (
	"log"

	"github.com/daqing/waterway/lib/utils"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var __gormDB__ *gorm.DB

func Setup() bool {
	dbType := utils.GetEnv("DB_TYPE")
	dbUrl := utils.GetEnv("DB_URL")

	if dbType == utils.EMPTY_STRING || dbUrl == utils.EMPTY_STRING {
		return false
	}

	var err error

	switch dbType {
	case "sqlite":
		__gormDB__, err = gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	case "pg":
		__gormDB__, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	default:
		log.Fatalf("unsupported database type: %s with db url: %s", dbType, dbUrl)
	}

	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

	return true
}

func DB() *gorm.DB {
	if __gormDB__ == nil {
		panic("database was not setup")
	}

	return __gormDB__
}
