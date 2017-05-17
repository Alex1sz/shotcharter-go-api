package main

import (
	"fmt"
	"github.com/alex1sz/shotcharter-go-api/db"
	"github.com/alex1sz/shotcharter-go-api/routers"
	"github.com/codegangsta/negroni"
	"github.com/unrolled/secure"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	db.Db.Ping()
	port := ":" + os.Getenv("PORT")
	router := routers.InitRoutes()

	secureMiddleware := secure.New(secure.Options{
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
	})
	n := negroni.Classic()

	n.Use(negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext))
	n.UseHandler(router)

	server := &http.Server{
		Handler: n,
		Addr:    port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Now listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}
