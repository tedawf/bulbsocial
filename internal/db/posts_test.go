package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomPost(t *testing.T) Post {
	user := CreateRandomTestUser(t)
	arg := CreatePostParams{
		Content: RandomContent(),
		Title:   RandomTitle(),
		UserID:  user.ID,
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
	CreateRandomPost(t)
}
