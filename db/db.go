package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)


type DB struct {
	Connection *sql.DB
}

func New(connection *sql.DB) *DB {
	return &DB{
		Connection: connection,
	}
}
