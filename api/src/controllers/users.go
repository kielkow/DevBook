package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateUser func
func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	user.ID, error = repository.Create(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// SearchUsers func
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching users"))
}

// SearchUser func
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching a user"))
}

// UpdateUser func
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating a user"))
}

// DeleteUser func
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a user"))
}
