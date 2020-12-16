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

// SearchByID publication by id
func (repository Publications) SearchByID(publicationID uint64) (models.Publication, error) {
	line, error := repository.db.Query(`
		SELECT p.*, u.nick FROM 
		publications p INNER JOIN users u 
		ON u.id = p.author_id WHERE p.id = ?`,
		publicationID,
	)

	if error != nil {
		return models.Publication{}, error
	}

	defer line.Close()

	var publication models.Publication
	if line.Next() {
		if error = line.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); error != nil {
			return models.Publication{}, error
		}
	}

	return publication, nil
}
