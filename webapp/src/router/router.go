package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate the app router
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Config(r)
}
