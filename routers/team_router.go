package routers

import (
	"github.com/alex1sz/shotcharter-go-api/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetTeamRoutes(router *mux.Router) *mux.Router {
	teamRouter := mux.NewRouter()
	teamRouter.HandleFunc("/teams", controllers.CreateTeam).Methods("POST")
	teamRouter.HandleFunc("/teams/{id}", controllers.GetTeamByID).Methods("GET")
	teamRouter.HandleFunc("/teams/{id}", controllers.UpdateTeam).Methods("PATCH")
	router.PathPrefix("/teams").Handler(negroni.New(
		negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext),
		negroni.Wrap(teamRouter),
		negroni.NewLogger(),
	))
	return router
}
