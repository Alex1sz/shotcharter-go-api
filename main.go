package main

import (
	"crypto/tls"
	// "fmt"
	"github.com/alex1sz/shotcharter-go-api/routers"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := ":" + os.Getenv("PORT")
	router := routers.InitRoutes()
	cert := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(os.Getenv("HOST")),
		Cache:      autocert.DirCache("certs"),
	}

	server := &http.Server{
		Handler:   router,
		Addr:      port,
		TLSConfig: &tls.Config{GetCertificate: cert.GetCertificate},
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// fmt.Printf("Now listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}
