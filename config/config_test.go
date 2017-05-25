package config_test

import (
	"github.com/alex1sz/shotcharter-go-api/config"
	"testing"
)

func TestGetConfig(t *testing.T) {
	var conf config.Config
	conf.GetConfig()

	if conf.DB.Driver != "postgres" || len(conf.DB.Connection) < 10 {
		t.Errorf("Expected conf.DB.Driver: postgres got: %v, & conf.DB.Connection str > 10 got: %v", conf.DB.Driver, conf.DB.Connection)
	}
}
