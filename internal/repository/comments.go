package repository

import (
	"context"

	"github.com/tedawf/bulbsocial/internal/db/sqlc"
)

type CommentRepository interface {
	GetCommentsByPostID(ctx context.Context, postID int64) ([]sqlc.GetCommentsByPostIDRow, error)
	CreateComment(ctx context.Context, params sqlc.CreateCommentParams) (sqlc.Comment, error)
}

type commentRepo struct {
	store *Store
}

func (c *commentRepo) CreateComment(ctx context.Context, params sqlc.CreateCommentParams) (sqlc.Comment, error) {
	return c.store.Queries.CreateComment(ctx, params)
}

func (c *commentRepo) GetCommentsByPostID(ctx context.Context, postID int64) ([]sqlc.GetCommentsByPostIDRow, error) {
	return c.store.Queries.GetCommentsByPostID(ctx, postID)

}

func NewCommentRepository(store *Store) CommentRepository {
	return &commentRepo{
		store: store,
	}
}
