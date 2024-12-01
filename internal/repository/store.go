package repository

import (
	"database/sql"

	"github.com/tedawf/bulbsocial/internal/db/sqlc"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*sqlc.Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: sqlc.New(db),
	}
}
