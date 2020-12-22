package controllers

import "net/http"

// RenderLoginScreen func
func RenderLoginScreen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Screen"))
}
