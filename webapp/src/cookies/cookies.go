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

// Read cookies
func Read(r *http.Request) (map[string]string, error) {
	cookie, error := r.Cookie("data")
	if error != nil {
		return nil, error
	}

	values := make(map[string]string)
	if error = s.Decode("data", cookie.Value, &values); error != nil {
		return nil, error
	}

	return values, nil
}
