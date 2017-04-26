package routers

import (
	"github.com/alex1sz/shotcharter-go/controllers"
	"github.com/gorilla/mux"
)

func SetTeamRoutes(router *mux.Router) *mux.Router {
	teamRouter := mux.NewRouter()
	subRouter := teamRouter.PathPrefix("/teams").Subrouter()
	subRouter.HandleFunc("/", controllers.CreateTeam).Methods("POST")

	subRouter.HandleFunc("/{id}", controllers.GetTeamByID).Methods("GET")

	return teamRouter
}
