package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomTestUser(t *testing.T) CreateUserRow {
	username := RandomUsername()
	arg := CreateUserParams{
		Username:       username,
		Email:          username + "@email.com",
		HashedPassword: []byte("123123"),
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
	CreateRandomTestUser(t)
}

func TestGetUserByID(t *testing.T) {
	user1 := CreateRandomTestUser(t)
	fmt.Print(user1.ID)
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
	user1 := CreateRandomTestUser(t)

	arg := UpdateUserPasswordParams{
		ID:             user1.ID,
		HashedPassword: []byte("4546456"),
	}

	err := testQueries.UpdateUserPassword(context.Background(), arg)
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	user1 := CreateRandomTestUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUserByID(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
