package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/tedawf/bulbsocial/internal/db"
)

var ErrUserNotFound = errors.New("user not found")

type PostService struct {
	store db.Store
}

func NewPostService(store db.Store) *PostService {
	return &PostService{store: store}
}

func (p *PostService) CreatePost(ctx context.Context, params db.CreatePostParams) (db.Post, error) {
	post, err := p.store.CreatePost(ctx, params)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation":
				return db.Post{}, ErrUserNotFound
			}
		}
		return db.Post{}, fmt.Errorf("unable to create post: %w", err)
	}
	return post, nil
}

func (p *PostService) GetPostByID(ctx context.Context, postID int64) (db.Post, error) {
	return p.store.GetPostByID(ctx, postID)
}

func (p *PostService) GetAllPosts(ctx context.Context, limit, offset int32) (posts []db.Post, err error) {
	return p.store.GetAllPosts(ctx, db.GetAllPostsParams{
		Limit:  limit,
		Offset: offset,
	})
}
