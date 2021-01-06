package routes

import (
	"net/http"
	"webapp/src/middlewares"

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
	routes = append(routes, routesUser...)
	routes = append(routes, routeHomePage)

	for _, route := range routes {
		if route.RequireAuthentication {
			router.HandleFunc(
				route.URI, 
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(
				route.URI, 
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
