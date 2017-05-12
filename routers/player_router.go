package routers

import (
	"github.com/alex1sz/shotcharter-go/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetPlayerRoutes(router *mux.Router) *mux.Router {
	playerRouter := mux.NewRouter()
	playerRouter.HandleFunc("/players", controllers.CreatePlayer).Methods("POST")
	router.PathPrefix("/players").Handler(negroni.New(negroni.Wrap(playerRouter)))

	return router
}
