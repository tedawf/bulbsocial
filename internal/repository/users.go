package repository

import (
	"context"
	"time"

	"github.com/tedawf/bulbsocial/internal/db/sqlc"
)

type UserRepository interface {
	GetUserByID(context.Context, int64) (sqlc.User, error)
	GetUserByEmail(context.Context, string) (sqlc.User, error)
	CreateUser(context.Context, sqlc.CreateUserParams) (sqlc.User, error)
	UpdateUser(context.Context, sqlc.UpdateUserParams) (sqlc.User, error)
	DeleteUser(context.Context, int64) error
	CreateAndInviteUser(ctx context.Context, params sqlc.CreateUserParams, token string, exp time.Duration) (sqlc.UserVerification, error)
	VerifyUser(ctx context.Context, token string) (sqlc.User, error)
}

type userRepo struct {
	store *Store
}

func (u *userRepo) CreateAndInviteUser(ctx context.Context, params sqlc.CreateUserParams, token string, exp time.Duration) (sqlc.UserVerification, error) {
	var invite sqlc.UserVerification
	err := u.store.execTx(ctx, func(q *sqlc.Queries) error {
		user, err := q.CreateUser(ctx, params)
		if err != nil {
			return err
		}

		invite, err = q.CreateUserVerification(ctx, sqlc.CreateUserVerificationParams{
			Token:  []byte(token),
			UserID: user.ID,
			Expiry: time.Now().Add(exp),
		})
		if err != nil {
			return err
		}

		return nil
	})
	return invite, err
}

func (u *userRepo) VerifyUser(ctx context.Context, token string) (sqlc.User, error) {
	var user sqlc.User

	err := u.store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error

		var userResult sqlc.GetUserFromInvitationRow
		userResult, err = q.GetUserFromInvitation(ctx, sqlc.GetUserFromInvitationParams{
			Token:  []byte(token),
			Expiry: time.Now(),
		})
		if err != nil {
			return err
		}

		userResult.IsVerified = true
		if user, err = u.store.UpdateUser(ctx, sqlc.UpdateUserParams{
			Username:   userResult.Username,
			Email:      userResult.Email,
			IsVerified: userResult.IsVerified,
			ID:         userResult.ID,
		}); err != nil {
			return err
		}

		if err = u.store.DeleteUserVerification(ctx, user.ID); err != nil {
			return err
		}

		return nil
	})

	return user, err
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

func NewUserRepository(store *Store) UserRepository {
	return &userRepo{
		store: store,
	}
}
