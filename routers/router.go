package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router = SetTeamRoutes(router)
	router = SetGameRoutes(router)

	return router
}
