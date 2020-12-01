package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users repository
type Users struct {
	db *sql.DB
}

// NewUsersRepository create a new users repository
func NewUsersRepository(db *sql.DB) *Users {
	return &users{db}
}

// Create a user on database
func (u Users) Create(user models.User) (uint64, error) {
	return 0, nil
}
