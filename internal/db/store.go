package db

import (
	"context"
	"database/sql"
	"errors"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	queries *Queries
	db      *sql.DB
}

// Creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		queries: New(db),
		db:      db,
	}
}

// ExecTx executes a function within a database transaction
func (s *Store) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	qtx := s.queries.WithTx(tx)

	if err = fn(qtx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.Join(err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
