package main

import (
	"fmt"
	"github.com/alex1sz/shotcharter-go-api/db"
	"github.com/alex1sz/shotcharter-go-api/routers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	db.Db.Ping()
	router := routers.InitRoutes()
	port := ":" + os.Getenv("PORT")

	server := &http.Server{
		Handler: router,
		Addr:    port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Now listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}
