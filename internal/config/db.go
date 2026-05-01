package config

import (
	"database/sql"

	_ "github.com/lib/pq" // Blank import to register the driver with init().
)

func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)

	return db, nil
}
