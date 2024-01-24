package repo

import (
	"errors"
	"log"

	"github.com/daqing/waterway/lib/utils"
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var __gormDB__ *gorm.DB

var ErrEmptyConfig = errors.New("database configuration is empty")
var ErrInvalidDatabaseType = errors.New("database type is invalid")
var ErrNotSetup = errors.New("database is not setup yet")

func Setup() error {
	viper.SetConfigFile("./config/db.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	dbType := viper.GetString("db_type")
	dbUrl := viper.GetString("db_url")

	if dbType == utils.EMPTY_STRING || dbUrl == utils.EMPTY_STRING {
		log.Println("db_type or db_url is empty")

		return ErrEmptyConfig
	}

	var err error

	switch dbType {
	case "sqlite":
		__gormDB__, err = gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	case "pg":
		__gormDB__, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	default:
		return ErrInvalidDatabaseType
	}

	if err != nil {
		log.Printf("failed to open database from gorm: %v", err)
		return err
	}

	log.Printf("Initialized database with type: %s and url: %s\n", dbType, dbUrl)

	return nil
}

func DB() (*gorm.DB, error) {
	if __gormDB__ == nil {
		return nil, ErrNotSetup
	}

	return __gormDB__, nil
}
