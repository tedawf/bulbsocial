package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tedawf/bulbsocial/internal/auth"
	"github.com/tedawf/bulbsocial/internal/util"
)

func createRandomTestUser(t *testing.T) User {
	username := util.RandomUsername()
	hashedPassword, err := auth.HashPassword(util.RandomString(10))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       username,
		Email:          username + "@email.com",
		HashedPassword: []byte(hashedPassword),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomTestUser(t)
}

func TestGetUserByID(t *testing.T) {
	user1 := createRandomTestUser(t)
	user2, err := testQueries.GetUserByID(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUserPassword(t *testing.T) {
	user1 := createRandomTestUser(t)
	hashedPassword, err := auth.HashPassword(util.RandomString(10))
	require.NoError(t, err)

	arg := UpdateUserPasswordParams{
		ID:             user1.ID,
		HashedPassword: []byte(hashedPassword),
	}

	err = testQueries.UpdateUserPassword(context.Background(), arg)
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomTestUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUserByID(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
