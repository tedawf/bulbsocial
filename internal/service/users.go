package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/tedawf/bulbsocial/internal/auth"
	"github.com/tedawf/bulbsocial/internal/db"
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
)

type UserService struct {
	store      db.Store
	tokenMaker auth.TokenMaker
}

func NewUserService(store db.Store, tokenMaker auth.TokenMaker) *UserService {
	return &UserService{store: store, tokenMaker: tokenMaker}
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

func (u *UserService) LoginUser(ctx context.Context, username, password string, duration time.Duration) (db.User, string, error) {
	user, err := u.store.GetUserByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return db.User{}, "", ErrInvalidCredentials
		}
		return db.User{}, "", fmt.Errorf("failed to fetch user: %w", err)
	}

	err = auth.CheckPassword(string(user.HashedPassword), password)
	if err != nil {
		return db.User{}, "", ErrInvalidCredentials
	}

	accessToken, err := u.tokenMaker.CreateToken(user.ID, duration)
	if err != nil {
		return db.User{}, "", fmt.Errorf("failed to create token: %w", err)
	}

	return user, accessToken, nil
}
