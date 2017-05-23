package db

import (
	"github.com/alex1sz/shotcharter-go-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	// "log"
)

var Db *sqlx.DB

func init() {
	var conf config.Config
	conf.GetConfig()

	Db = sqlx.MustConnect("postgres", conf.DB.Connection)
	// sanity check values before deploying production
	Db.SetMaxIdleConns(4)
	Db.SetMaxOpenConns(16)
}
