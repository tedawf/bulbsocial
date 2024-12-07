package service

import (
	"context"

	"github.com/tedawf/bulbsocial/internal/auth"
	"github.com/tedawf/bulbsocial/internal/db"
)

type UserService struct {
	store db.Store
}

func NewUserService(store db.Store) *UserService {
	return &UserService{store: store}
}

func (u *UserService) GetUserByID(ctx context.Context, userID int64) (db.User, error) {
	return u.store.GetUserByID(ctx, userID)
}

func (u *UserService) CreateUser(ctx context.Context, username, email, password string) (db.User, error) {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return db.User{}, err
	}

	params := db.CreateUserParams{
		Username:       username,
		Email:          email,
		HashedPassword: []byte(hashedPassword),
	}

	return u.store.CreateUser(ctx, params)
}

func (u *UserService) DeleteUser(ctx context.Context, userID int64) error {
	return u.store.DeleteUser(ctx, userID)
}
