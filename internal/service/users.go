package service

import (
	"context"

	"github.com/tedawf/bulbsocial/internal/db"
)

type UserService struct {
	store db.Store
}

func NewUserService(store db.Store) *UserService {
	return &UserService{store: store}
}

func (u *UserService) GetUserByID(ctx context.Context, userID int64) (user db.User, err error) {
	return user, u.store.ExecTx(ctx, func(q db.Querier) error {
		user, err = q.GetUserByID(ctx, userID)
		return err
	})
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (user db.User, err error) {
	return user, u.store.ExecTx(ctx, func(q db.Querier) error {
		user, err = q.GetUserByEmail(ctx, email)
		return err
	})
}

func (u *UserService) CreateUser(ctx context.Context, username, email, password string) (user db.CreateUserRow, err error) {
	return user, u.store.ExecTx(ctx, func(q db.Querier) error {
		params := db.CreateUserParams{
			Username:       username,
			Email:          email,
			HashedPassword: []byte(password), // todo: hash
		}

		user, err = q.CreateUser(ctx, params)
		return err
	})
}

func (u *UserService) DeleteUser(ctx context.Context, userID int64) error {
	return u.store.ExecTx(ctx, func(q db.Querier) error {
		return q.DeleteUser(ctx, userID)
	})
}
