package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func setupDatabase(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		return nil, err
	}

	return db, nil
}
