package repository

import (
	"context"

	"github.com/tedawf/bulbsocial/internal/db/sqlc"
)

type RoleRepository interface {
	GetByName(ctx context.Context, name string) (sqlc.Role, error)
}

type roleRepo struct {
	store *Store
}

func (r *roleRepo) GetByName(ctx context.Context, name string) (sqlc.Role, error) {
	return r.store.Queries.GetRoleByName(ctx, name)
}

func NewRoleRepository(store *Store) RoleRepository {
	return &roleRepo{
		store: store,
	}
}
