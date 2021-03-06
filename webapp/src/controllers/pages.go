package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// RenderLoginScreen func
func RenderLoginScreen(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

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

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutingTemplate(w, "home.html", struct {
		Publications []models.Publication
		UserID       uint64
	}{
		Publications: publications,
		UserID:       userID,
	})
}

// RenderUpdatePublicationPage func
func RenderUpdatePublicationPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, error := strconv.ParseUint(params["publicationId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.APIURL, publicationID)
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

	var publication models.Publication
	if error = json.NewDecoder(response.Body).Decode(&publication); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: error.Error()})
		return
	}

	utils.ExecutingTemplate(w, "update-publication.html", publication)
}

// RenderUsersPage func
func RenderUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?user=%s", config.APIURL, nameOrNick)

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

	var users []models.User
	if error = json.NewDecoder(response.Body).Decode(&users); error != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: error.Error()})
		return
	}

	utils.ExecutingTemplate(w, "users.html", users)
}

// RenderUserProfile func
func RenderUserProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	signinUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userID == signinUserID {
		http.Redirect(w, r, "/profile", 302)
		return
	}

	user, error := models.SearchCompletedUser(userID, r)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}

	utils.ExecutingTemplate(w, "user.html", struct {
		User         models.User
		SigninUserID uint64
	}{
		User:         user,
		SigninUserID: signinUserID,
	})
}

// RenderSigninUserProfile func
func RenderSigninUserProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, error := models.SearchCompletedUser(userID, r)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}

	utils.ExecutingTemplate(w, "profile.html", user)
}

// RenderEditUserPage func
func RenderEditUserPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channel := make(chan models.User)
	go models.SearchUserData(channel, userID, r)
	user := <-channel

	if user.ID == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Error to search user"})
		return
	}

	utils.ExecutingTemplate(w, "edit-user.html", user)
}

// RenderUpdatePasswordPage func
func RenderUpdatePasswordPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecutingTemplate(w, "update-password.html", nil)
}
