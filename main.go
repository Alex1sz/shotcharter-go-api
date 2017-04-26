package main

import (
	"github.com/alex1sz/shotcharter-go/db"
	"github.com/alex1sz/shotcharter-go/routers"
	"log"
	"net/http"
	"time"
)

func main() {
	db.Db.Ping()
	router := routers.InitRoutes()

	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Now listening on port: 127.0.0.1:8080")
	log.Fatal(server.ListenAndServe())
}
