// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (CreatePostRow, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	DeleteComment(ctx context.Context, id int64) error
	DeletePost(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	FollowUser(ctx context.Context, arg FollowUserParams) error
	GetAllPosts(ctx context.Context, arg GetAllPostsParams) ([]Post, error)
	GetCommentsByPost(ctx context.Context, arg GetCommentsByPostParams) ([]Comment, error)
	GetFollowees(ctx context.Context, arg GetFolloweesParams) ([]int64, error)
	GetFollowers(ctx context.Context, arg GetFollowersParams) ([]int64, error)
	GetPostByID(ctx context.Context, id int64) (Post, error)
	GetPostsByUser(ctx context.Context, arg GetPostsByUserParams) ([]Post, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	SearchComments(ctx context.Context, arg SearchCommentsParams) ([]Comment, error)
	SearchPosts(ctx context.Context, arg SearchPostsParams) ([]Post, error)
	SearchUsers(ctx context.Context, arg SearchUsersParams) ([]SearchUsersRow, error)
	UnfollowUser(ctx context.Context, arg UnfollowUserParams) error
	UpdatePost(ctx context.Context, arg UpdatePostParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
