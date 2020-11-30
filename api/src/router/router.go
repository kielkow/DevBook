package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate API Router
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Config(r)
}
