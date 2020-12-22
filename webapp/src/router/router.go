package router

import "github.com/gorilla/mux"

// Generate the app router
func Generate() *mux.Router {
	return mux.NewRouter()
}
