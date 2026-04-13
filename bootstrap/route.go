package bootstrap

import (
	"github.com/gorilla/mux"
	"github.com/merkurtran/goblog/routes"
)

func SetRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	return router
}
