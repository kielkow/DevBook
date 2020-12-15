package repositories

import (
	"api/src/models"
	"database/sql"
)

// Publications repository
type Publications struct {
	db *sql.DB
}

// NewPublicationsRepository create a new publications repository
func NewPublicationsRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

// Create a publication on database
func (repository Publications) Create(publication models.Publication) (uint64, error) {
	statement, error := repository.db.Prepare("INSERT INTO publications (title, content, author_id) VALUES (?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if error != nil {
		return 0, error
	}

	lastIDInserted, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastIDInserted), nil
}
