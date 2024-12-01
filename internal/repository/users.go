package repository

import (
	"context"

	"github.com/tedawf/bulbsocial/internal/db/sqlc"
)

type UserRepository interface {
	GetUserByID(context.Context, int64) (sqlc.User, error)
	GetUserByEmail(context.Context, string) (sqlc.User, error)
	CreateUser(context.Context, sqlc.CreateUserParams) (sqlc.User, error)
	UpdateUser(context.Context, sqlc.UpdateUserParams) (sqlc.User, error)
	DeleteUser(context.Context, int64) error
}

type userRepo struct {
	store *Store
}

func NewUserRepository(store *Store) UserRepository {
	return &userRepo{
		store: store,
	}
}

func (u *userRepo) GetUserByID(ctx context.Context, id int64) (sqlc.User, error) {
	return u.store.Queries.GetUserByID(ctx, id)
}

func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (sqlc.User, error) {
	return u.store.Queries.GetUserByEmail(ctx, email)
}

func (u *userRepo) CreateUser(ctx context.Context, params sqlc.CreateUserParams) (sqlc.User, error) {
	return u.store.Queries.CreateUser(ctx, params)
}

func (u *userRepo) UpdateUser(ctx context.Context, params sqlc.UpdateUserParams) (sqlc.User, error) {
	return u.store.Queries.UpdateUser(ctx, params)
}

func (u *userRepo) DeleteUser(ctx context.Context, id int64) error {
	return u.store.Queries.DeleteUser(ctx, id)
}
