package config

type Config struct {
	Db struct {
		Connection string `default:"dbname=shotcharter_go_development host=localhost sslmode=disable" env:"CONFIG_DB_CONNECTION"`
		Driver     string `default:"postgres"`
	}
}
