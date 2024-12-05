package db

import (
	"context"
	"database/sql"
	"errors"
)

type Store interface {
	Querier
	ExecTx(ctx context.Context, fn func(Querier) error) error
}

// SQLStore provides all functions to execute sql queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// Creates a new SQLStore
func NewSQLStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}

// ExecTx executes a function within a database transaction
func (s *SQLStore) ExecTx(ctx context.Context, fn func(Querier) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	qtx := s.Queries.WithTx(tx)

	if err = fn(qtx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.Join(err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
