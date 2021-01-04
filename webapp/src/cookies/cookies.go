package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Config cookies
func Config() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Save cookie
func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encodingData, error := s.Encode("data", data)
	if error != nil {
		return error
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encodingData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
