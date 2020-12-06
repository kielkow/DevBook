package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User struct
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

// Prepare user model
func (user *User) Prepare(step string) error {
	if error := user.validate(step); error != nil {
		return error
	}

	user.format()

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Name is required")
	}

	if user.Nick == "" {
		return errors.New("Nick is required")
	}

	if user.Email == "" {
		return errors.New("E-mail is required")
	}

	if error := checkmail.ValidateFormat(user.Email); error != nil {
		return errors.New("E-mail invalid")
	}

	if step == "create" && user.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
