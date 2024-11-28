package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	username := RandomUser()
	arg := CreateUserParams{
		Username: username,
		Email:    username + "@email.com",
		Password: []byte("123123"),
		Name:     "user",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}
