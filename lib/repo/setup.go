package repo

import (
	"log"

	"github.com/daqing/waterway/lib/utils"
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var __gormDB__ *gorm.DB

func Setup() bool {
	viper.SetConfigFile("./config/db.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("can't read config file: %v", err)
	}

	dbType := viper.GetString("db_type")
	dbUrl := viper.GetString("db_url")

	if dbType == utils.EMPTY_STRING || dbUrl == utils.EMPTY_STRING {
		log.Println("db_type or db_url is empty")

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

	log.Printf("Initialized database with type: %s and url: %s\n", dbType, dbUrl)

	return true
}

func DB() *gorm.DB {
	if __gormDB__ == nil {
		panic("database was not setup")
	}

	return __gormDB__
}
