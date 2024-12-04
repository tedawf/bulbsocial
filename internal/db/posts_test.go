package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomTestPost(t *testing.T) Post {
	arg := CreatePostParams{
		Content: RandomContent(),
		Title:   RandomTitle(),
		UserID:  CreateRandomTestUser(t).ID,
		Tags:    RandomTags(),
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.Content, post.Content)
	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.UserID, post.UserID)
	require.Equal(t, arg.Tags, post.Tags)

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
		Title:   RandomTitle(),
		Content: RandomContent(),
		ID:      post1.ID,
		Version: post1.Version,
	}

	version, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, version)

	require.NotEqual(t, arg.Version, version)
}
