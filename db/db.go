package db

import (
	"database/sql"

	"github.com/cauabernardino/go-rest-api/config"

	_ "github.com/lib/pq"
)

// Connect opens the connection with Database
func Connect() (*sql.DB, error) {

	db, err := sql.Open("postgres", config.DBConnectString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
