package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Signout func
func Signout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", 302)
}
