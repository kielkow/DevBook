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
	return &Users{db}
}

// Create a user on database
func (repository Users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastIDInserted, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastIDInserted), nil
}
