package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	SetTeamRoutes(router)
	SetGameRoutes(router)
	SetPlayerRoutes(router)

	return router
}
