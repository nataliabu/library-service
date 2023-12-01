package database

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const defaultTimeout = 3 * time.Second

type DB struct {
	*sqlx.DB
}

func New(dsn string) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "postgres", "postgres://"+dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)

	createTables(db)
	return &DB{db}, nil
}

func createTables(db *sqlx.DB) {
	queryBooks := `CREATE TABLE IF NOT EXISTS books (
    id serial PRIMARY KEY,
    title text NOT NULL,
    author text NOT NULL,
    isbn text NOT NULL,
    issue_year int NOT NULL,
		available boolean NOT NULL default true
	);`

	_, err := db.Exec(queryBooks)
	if err != nil {
		log.Fatal(err)
	}

	queryCustomers := `CREATE TABLE IF NOT EXISTS customers (
    id serial PRIMARY KEY,
    name text NOT NULL
	);`

	_, err = db.Exec(queryCustomers)
	if err != nil {
		log.Fatal(err)
	}
}
