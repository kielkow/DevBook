package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorAPI interface
type ErrorAPI struct {
	Error string `json:"error"`
}

// JSON response func
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if data != nil {
		if error := json.NewEncoder(w).Encode(data); error != nil {
			log.Fatal(error)
		}
	}
}

// TreatError response func
func TreatError(w http.ResponseWriter, r *http.Response) {
	var error ErrorAPI

	json.NewDecoder(r.Body).Decode(&error)

	JSON(w, r.StatusCode, error)
}
