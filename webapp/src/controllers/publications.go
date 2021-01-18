package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

// CreatePublication func
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publication, error := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications", config.APIURL)
	response, error := requests.DoAuthenticateRequest(r, http.MethodPost, url, bytes.NewBuffer(publication))
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// LikePublication func
func LikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, error := strconv.ParseUint(params["publicationId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d/like", config.APIURL, publicationID)
	response, error := requests.DoAuthenticateRequest(r, http.MethodPost, url, nil)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// DislikePublication func
func DislikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, error := strconv.ParseUint(params["publicationId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d/dislike", config.APIURL, publicationID)
	response, error := requests.DoAuthenticateRequest(r, http.MethodPost, url, nil)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// UpdatePublication func
func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, error := strconv.ParseUint(params["publicationId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	r.ParseForm()
	publication, error := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.APIURL, publicationID)
	response, error := requests.DoAuthenticateRequest(r, http.MethodPut, url, bytes.NewBuffer(publication))
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// DeletePublication func
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, error := strconv.ParseUint(params["publicationId"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.APIURL, publicationID)
	response, error := requests.DoAuthenticateRequest(r, http.MethodDelete, url, nil)
	if error != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
