package router

import "github.com/gorilla/mux"

// Generate API Router
func Generate() *mux.Router {
	return mux.NewRouter()
}
