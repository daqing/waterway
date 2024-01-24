package repo

import (
	"github.com/daqing/waterway/lib/clients"
	"github.com/daqing/waterway/lib/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Setup() bool {
	dbType := utils.GetEnv("DB_TYPE")
	dbUrl := utils.GetEnv("DB_URL")

	if dbType == utils.EMPTY_STRING || dbUrl == utils.EMPTY_STRING {
		return false
	}

	Pool = clients.NewPGPool(utils.GetEnvMust("WATERWAY_PG_URL"))

	return true
}
