package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Connect opens the database connection and returns it
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", ConectionDB)
	if err != nil {
		log.Fatal("could not open connect with databse: %w", err)

		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatal("could not close port database: %w", err)

		return nil, err
	}

	return db, nil
}
