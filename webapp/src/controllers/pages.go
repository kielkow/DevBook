package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// RenderLoginScreen func
func RenderLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecutingTemplate(w, "login.html", nil)
}

// RenderSignupScreen func
func RenderSignupScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecutingTemplate(w, "signup.html", nil)
}
