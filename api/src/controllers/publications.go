package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreatePublication func
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var publication models.Publication
	if error = json.Unmarshal(requestBody, &publication); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	publication.AuthorID = userID

	if error = publication.Prepare(); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)
	publication.ID, error = repository.Create(publication)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}

// SearchPublications func
func SearchPublications(w http.ResponseWriter, r *http.Request) {

}

// SearchPublication func
func SearchPublication(w http.ResponseWriter, r *http.Request) {

}

// UpdatePublication func
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}

// DeletePublication func
func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
