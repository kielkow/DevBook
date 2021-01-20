package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger write request info on terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Authenticate check if cookies exists
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, error := cookies.Read(r); error != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		nextFunction(w, r)
	}
}
