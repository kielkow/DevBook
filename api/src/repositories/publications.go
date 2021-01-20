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

// Search publications by user id from user and followers
func (repository Publications) Search(userID uint64) ([]models.Publication, error) {
	lines, error := repository.db.Query(`
		SELECT DISTINCT p.*, u.nick FROM publications p 
		INNER JOIN users u ON u.id = p.author_id
		INNER JOIN followers f ON p.author_id = f.user_id 
		WHERE u.id = ? OR f.follower_id = ?
		ORDER BY 1 desc`,
		userID, userID,
	)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var publications []models.Publication
	for lines.Next() {
		var publication models.Publication

		if error = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); error != nil {
			return nil, error
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Update publication by id
func (repository Publications) Update(publicationID uint64, publication models.Publication) error {
	statement, error := repository.db.Prepare("UPDATE publications SET title = ?, content = ? WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(publication.Title, publication.Content, publicationID); error != nil {
		return error
	}

	return nil
}

// Delete user by id
func (repository Publications) Delete(publicationID uint64) error {
	statement, error := repository.db.Prepare("DELETE FROM publications WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(publicationID); error != nil {
		return error
	}

	return nil
}

// SearchByUserID func
func (repository Publications) SearchByUserID(userID uint64) ([]models.Publication, error) {
	lines, error := repository.db.Query(`
		SELECT p.*, u.nick FROM publications p 
		JOIN users u ON u.id = p.author_id
		WHERE p.author_id = ?`,
		userID,
	)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var publications []models.Publication
	for lines.Next() {
		var publication models.Publication

		if error = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); error != nil {
			return nil, error
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Like func
func (repository Publications) Like(publicationID uint64) error {
	statement, error := repository.db.Prepare("UPDATE publications SET likes = likes + 1 WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(publicationID); error != nil {
		return error
	}

	return nil
}

// Unlike func
func (repository Publications) Unlike(publicationID uint64) error {
	statement, error := repository.db.Prepare(`
		UPDATE publications SET likes =
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE 0 
		END
		WHERE id = ?
	`)
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(publicationID); error != nil {
		return error
	}

	return nil
}
