package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
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

// RenderHomePage func
func RenderHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.APIURL)
	response, error := requests.DoAuthenticateRequest(r, http.MethodGet, url, nil)
	fmt.Println(response.StatusCode, error)

	utils.ExecutingTemplate(w, "home.html", nil)
}
