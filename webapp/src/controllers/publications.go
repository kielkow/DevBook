package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"
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
