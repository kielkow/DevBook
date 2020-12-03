package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Search users by name or nick on database
func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, error := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? OR nick LIKE ?",
		nameOrNick, nameOrNick)

	if error != nil {
		return nil, error
	}

	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if error := lines.Scan(
			&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

// SearchByID user by id
func (repository Users) SearchByID(ID uint64) (models.User, error) {
	line, error := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE id = ?",
		ID)

	if error != nil {
		return models.User{}, error
	}

	defer line.Close()

	var user models.User
	if line.Next() {
		if error = line.Scan(
			&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

// Update user by id
func (repository Users) Update(ID uint64, user models.User) error {
	statement, error := repository.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(user.Name, user.Nick, user.Email, ID); error != nil {
		return error
	}

	return nil
}
