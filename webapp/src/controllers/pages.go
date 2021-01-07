package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
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
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	var publications []models.Publication
	if error = json.NewDecoder(response.Body).Decode(&publications); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: error.Error()})
		return
	}

	utils.ExecutingTemplate(w, "home.html", publications)
}
