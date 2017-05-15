package config

type Config struct {
	Db struct {
		Connection string
		Driver     string
	}
}
