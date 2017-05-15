package config_test

import (
	"github.com/alex1sz/configor"
	"github.com/alex1sz/shotcharter-go-api/config"
	"testing"
)

func TestConfigInit(t *testing.T) {
	var config config.Config
	configor.Load(&config, "test.db_conf.yml")

	if config.Db.Connection != "dbname=shotcharter_go_test host=localhost sslmode=disable" {
		t.Error("config package init() Connection failed!")
	}

	if config.Db.Driver != "postgres" {
		t.Error("config package init() failed due to DriverName!")
	}
}
