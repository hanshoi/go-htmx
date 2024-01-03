package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type DB struct {
	db    *sql.DB
	Lists ListRepo
}

func New() DB {
	conn, err := sql.Open("pgx", "postgresql://admin:admin@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	if err = conn.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}

	return DB{
		Lists: ListRepo{conn},
	}
}

func (db *DB) Close() {
	db.Close()
}
