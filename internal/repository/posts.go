package repository

import (
	"context"
	"database/sql"

	"github.com/tedawf/bulbsocial/internal/db/sqlc"
)

type PostRepository interface {
	GetPostByID(context.Context, int64) (sqlc.Post, error)
	CreatePost(context.Context, sqlc.CreatePostParams) (sqlc.Post, error)
	Update(context.Context, sqlc.UpdatePostParams) (sql.NullInt32, error)
	Delete(context.Context, int64) error
}

type postRepo struct {
	store *Store
}

func NewPostRepository(store *Store) PostRepository {
	return &postRepo{
		store: store,
	}
}

func (p *postRepo) GetPostByID(ctx context.Context, id int64) (sqlc.Post, error) {
	return p.store.Queries.GetPostByID(ctx, id)
}

func (p *postRepo) CreatePost(ctx context.Context, params sqlc.CreatePostParams) (sqlc.Post, error) {
	return p.store.Queries.CreatePost(ctx, params)
}

func (p *postRepo) Update(ctx context.Context, params sqlc.UpdatePostParams) (sql.NullInt32, error) {
	return p.store.UpdatePost(ctx, params)
}

func (p *postRepo) Delete(ctx context.Context, id int64) error {
	return p.store.DeletePost(ctx, id)
}
