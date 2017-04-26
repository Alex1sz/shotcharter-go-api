package routers

import (
	"github.com/alex1sz/shotcharter-go/controllers"
	"github.com/gorilla/mux"
)

func SetGameRoutes(router *mux.Router) *mux.Router {
	gameRouter := mux.NewRouter()
	subRouter := gameRouter.PathPrefix("/games").Subrouter()

	subRouter.HandleFunc("/{id}", controllers.GetGameByID).Methods("GET")

	return router
}
