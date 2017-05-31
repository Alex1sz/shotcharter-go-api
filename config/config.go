package config

import (
	"os"
	"regexp"
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

func matchTestStr() string {
	isTest, _ := regexp.MatchString("/_test/", os.Args[0])

	if isTest {
		return "test"
	}
	return ""
}

// GetEnvironment returns configor.Environment str
func (config *Config) getEnvironment() {
	if config.Environment != "" {
		return
	}
	isTestStr := matchTestStr()
	envStrings := []string{os.Getenv("CONFIGOR_ENV"), isTestStr, "development"}

	for _, envStr := range envStrings {
		if envStr != "" {
			config.Environment = envStr
			break
		}
	}
	return
}

func (conf *Config) GetConfig() {
	if len(conf.Environment) < 1 {
		conf.getEnvironment()
	}

	conf.DB.Driver = "postgres"
	if len(conf.DB.Connection) < 5 {
		conf.setDbConnectionStr()
		return
	}
	return
}
