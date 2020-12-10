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

// Delete user by id
func (repository Users) Delete(ID uint64) error {
	statement, error := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(ID); error != nil {
		return error
	}

	return nil
}

// SearchByEmail user by e-mail
func (repository Users) SearchByEmail(email string) (models.User, error) {
	line, error := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if error != nil {
		return models.User{}, error
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if error = line.Scan(&user.ID, &user.Password); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

// Follow a user
func (repository Users) Follow(userID, followerID uint64) error {
	statement, error := repository.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userID, followerID); error != nil {
		return error
	}

	return nil
}

// Unfollow a user
func (repository Users) Unfollow(userID, followerID uint64) error {
	statement, error := repository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userID, followerID); error != nil {
		return error
	}

	return nil
}

// SearchFollowers from a user
func (repository Users) SearchFollowers(userID uint64) ([]models.User, error) {
	lines, error := repository.db.Query(`
		SELECT u.id, u.nome, u.nick, u.email, u.createdAt
		FROM users u INNER JOIN followers f ON u.id = f.follower_id WHERE f.user_id = ?`, 
		userID)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if error = lines.Scan(
			&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}
