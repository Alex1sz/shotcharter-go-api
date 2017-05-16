package config

import (
	"github.com/alex1sz/configor"
	// "log"
)

type Config struct {
	Db struct {
		Connection string
		Driver     string
	}
}

func (config Config) getDBConfFileStr() string {
	envStr := configor.ENV()
	defaultPath := "../config/"
	defaultFile := "db_conf.yml"

	if envStr == "development" {
		return defaultPath + defaultFile
	}

	if len(envStr) == 0 {
		return defaultPath + defaultFile
	}
	return defaultPath + envStr + "." + defaultFile
}

func SetConfig() Config {
	config := Config{}
	configor.Load(&config, config.getDBConfFileStr())
	return config
}
