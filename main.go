package main

import (
	"github.com/alex1sz/shotcharter-go/db"
)

func main() {
	db.Db.Ping()
}
