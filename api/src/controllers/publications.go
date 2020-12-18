package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	userID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)
	publications, error := repository.Search(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}

// SearchPublication func
func SearchPublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationID, error := strconv.ParseUint(params["publicationId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)

	publication, error := repository.SearchByID(publicationID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, publication)
}

// UpdatePublication func
func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)
	publicationID, error := strconv.ParseUint(params["publicationId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)
	publicationSaveOnDatabase, error := repository.SearchByID(publicationID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if publicationSaveOnDatabase.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("Not possible update a publcation that is not yours"))
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

	if error = publication.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = repository.Update(publicationID, publication); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletePublication func
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)
	publicationID, error := strconv.ParseUint(params["publicationId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)
	publicationSaveOnDatabase, error := repository.SearchByID(publicationID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if publicationSaveOnDatabase.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("Not possible delete a publication that is not yours"))
		return
	}

	if error = repository.Delete(publicationID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// SearchUserPublications func
func SearchUserPublications(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, error := strconv.ParseUint(params["userId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)
	publications, error := repository.SearchByUserID(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}
