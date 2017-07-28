package routers

import (
	"github.com/alex1sz/shotcharter-go-api/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetPlayerRoutes(router *mux.Router) *mux.Router {
	playerRouter := mux.NewRouter()
	playerRouter.HandleFunc("/players", controllers.CreatePlayer).Methods("POST")
	playerRouter.HandleFunc("/players/{id}", controllers.UpdatePlayer).Methods("PATCH")
	router.PathPrefix("/players").Handler(negroni.New(
		negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext),
		negroni.Wrap(playerRouter),
		negroni.NewLogger(),
	))
	return router
}
