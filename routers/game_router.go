package routers

import (
	"github.com/alex1sz/shotcharter-go/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetGameRoutes(router *mux.Router) *mux.Router {
	gameRouter := mux.NewRouter()
	gameRouter.HandleFunc("/games/{id}", controllers.GetGameByID).Methods("GET")
	router.PathPrefix("/games").Handler(negroni.New(negroni.Wrap(gameRouter)))

	return router
}
