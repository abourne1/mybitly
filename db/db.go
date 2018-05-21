package db

import (
	"sync"
	"database/sql"

	_ "github.com/lib/pq"
)


type DB struct {
	mu sync.Mutex
	Connection *sql.DB
}

func New(connection *sql.DB) *DB {
	return &DB{
		Connection: connection,
	}
}
