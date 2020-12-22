package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct
type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

// Config Router func
func Config(router *mux.Router) *mux.Router {
	routes := routesLogin

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return router
}
