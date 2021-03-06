package routers

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/secure"
)

var secureMiddleware = secure.New(secure.Options{
	// SSLRedirect:        true,
	FrameDeny:          true,
	ContentTypeNosniff: true,
	BrowserXssFilter:   true,
	IsDevelopment:      true,
})

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	SetTeamRoutes(router)
	SetGameRoutes(router)
	SetPlayerRoutes(router)
	SetShotRoutes(router)

	return router
}
