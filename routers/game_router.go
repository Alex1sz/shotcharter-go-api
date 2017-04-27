package routers

import (
	"github.com/alex1sz/shotcharter-go/controllers"
	"github.com/gorilla/mux"
)

func SetGameRoutes(router *mux.Router) *mux.Router {
	gameRouter := mux.NewRouter()
	gameRouter.HandleFunc("/games/{id}", controllers.GetGameByID).Methods("GET")

	return gameRouter
}
