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
func Config(r *mux.Router) *mux.Router {
	routes := routesUsers
	routes = append(routes, routeLogin)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
