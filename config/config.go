package config

import (
	"github.com/alex1sz/configor"
	"os"
)

type Config struct {
	Environment string

	DB struct {
		Connection string
		Driver     string
	}
}

func (conf *Config) setDbConnectionStr() {
	switch conf.Environment {
	case "test":
		conf.DB.Connection = "dbname=shotcharter_go_test host=localhost sslmode=disable"
	case "production":
		conf.DB.Connection = os.Getenv("DB_CONNECTION")
	default:
		conf.DB.Connection = "dbname=shotcharter_go_development host=localhost sslmode=disable"
	}
	return
}

func (conf *Config) GetConfig() {
	conf.DB.Driver = "postgres"
	if len(conf.DB.Connection) < 5 {
		conf.Environment = configor.ENV()
		conf.setDbConnectionStr()
		return
	}
	return
}
