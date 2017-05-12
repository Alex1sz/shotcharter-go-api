package routers

import (
	"github.com/alex1sz/shotcharter-go/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetShotRoutes(router *mux.Router) *mux.Router {
	shotRouter := mux.NewRouter()
	shotRouter.HandleFunc("/shots", controllers.CreateShot).Methods("POST")
	router.PathPrefix("/shots").Handler(negroni.New(negroni.Wrap(shotRouter)))

	return router
}
