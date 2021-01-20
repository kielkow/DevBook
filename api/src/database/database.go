package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver to connect with MySQL database
)

// Connect with database
func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.StringDatabaseConnection)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()

		return nil, error
	}

	return db, nil
}
