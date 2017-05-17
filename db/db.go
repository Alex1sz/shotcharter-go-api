package db

import (
	"github.com/alex1sz/configor"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	// "log"
)

func getDbConnectionStr() string {
	if configor.ENV() == "test" {
		return "dbname=shotcharter_go_test host=localhost sslmode=disable"
	}
	return "dbname=shotcharter_go_development host=localhost sslmode=disable"
}

var Db *sqlx.DB

func init() {
	Db = sqlx.MustConnect("postgres", getDbConnectionStr())
	// sanity check values before deploying production
	Db.SetMaxIdleConns(4)
	Db.SetMaxOpenConns(16)
}
