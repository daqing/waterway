package repo

import (
	"github.com/daqing/waterway/lib/clients"
	"github.com/daqing/waterway/lib/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Setup() {
	Pool = clients.NewPGPool(utils.GetEnvMust("WATERWAY_PG_URL"))
}
