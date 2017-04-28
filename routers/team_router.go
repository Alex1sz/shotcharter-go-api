package routers

import (
	"github.com/alex1sz/shotcharter-go/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetTeamRoutes(router *mux.Router) *mux.Router {
	teamRouter := mux.NewRouter()
	teamRouter.HandleFunc("/teams", controllers.CreateTeam).Methods("POST")
	teamRouter.HandleFunc("/teams/{id}", controllers.GetTeamByID).Methods("GET")
	router.PathPrefix("/teams").Handler(negroni.New(negroni.Wrap(teamRouter)))

	return router
}
