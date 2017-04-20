package main

import (
	"github.com/alex1sz/shotcharter-go/db"
	_ "github.com/lib/pq"
)

func main() {
	// fmt.Printf("Hello, world.\n")
	db.Init()
}
