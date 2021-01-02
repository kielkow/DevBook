package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Login to authenticate a user
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userSaved, error := repository.SearchByEmail(user.Email)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.ValidatePassword(userSaved.Password, user.Password); error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	token, error := authentication.CreateToken(userSaved.ID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	userID := strconv.FormatUint(userSaved.ID, 10)

	responses.JSON(w, http.StatusOK, models.AuthenticationData{ID: userID, Token: token})
}
