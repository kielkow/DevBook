package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser func
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		log.Fatal(error)
	}

	db, error := database.Connect()
	if error != nil {
		log.Fatal(error)
	}

	repository := repositories.NewUsersRepository(db)

	userID, error := repository.Create(user)
	if error != nil {
		log.Fatal(error)
	}

	w.Write([]byte(fmt.Sprintf("ID inserted: %d", userID)))
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
