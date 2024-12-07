package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tedawf/bulbsocial/internal/util"
)

func CreateRandomTestPost(t *testing.T) Post {
	arg := CreatePostParams{
		UserID:  createRandomTestUser(t).ID,
		Title:   util.RandomTitle(),
		Content: util.RandomContent(),
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.UserID, post.UserID)
	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.Content, post.Content)

	return post
}

func TestCreatePost(t *testing.T) {
	CreateRandomTestPost(t)
}

func TestDeletePost(t *testing.T) {
	post1 := CreateRandomTestPost(t)
	err := testQueries.DeletePost(context.Background(), post1.ID)
	require.NoError(t, err)

	post2, err := testQueries.GetPostByID(context.Background(), post1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post2)
}

func TestUpdatePost(t *testing.T) {
	post1 := CreateRandomTestPost(t)

	arg := UpdatePostParams{
		ID:      post1.ID,
		Title:   util.RandomTitle(),
		Content: util.RandomContent(),
	}

	err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
}
