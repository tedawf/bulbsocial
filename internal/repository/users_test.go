package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tedawf/bulbsocial/internal/db/sqlc"
)

func CreateRandomUserVerification(t *testing.T) sqlc.UserVerification {
	userRepo := NewUserRepository(NewStore(TestDB))

	username := sqlc.RandomUsername()
	arg := sqlc.CreateUserParams{
		Username: username,
		Email:    username + "@email.com",
		Password: []byte("123123"),
		Name:     "user",
	}
	token := sqlc.RandomUsername() + sqlc.RandomUsername()
	exp := time.Minute * 5

	invite, err := userRepo.CreateAndInviteUser(context.Background(), arg, token, exp)
	require.NoError(t, err)
	require.NotEmpty(t, invite)

	require.Equal(t, string(invite.Token), token)

	return invite
}

func TestCreateAndInviteUser(t *testing.T) {
	CreateRandomUserVerification(t)
}

func TestVerifyUser(t *testing.T) {
	userRepo := NewUserRepository(NewStore(TestDB))

	invite := CreateRandomUserVerification(t)

	user, err := userRepo.VerifyUser(context.Background(), string(invite.Token))
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.True(t, user.IsVerified)
}
