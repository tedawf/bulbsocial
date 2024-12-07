package service

import (
	"context"

	"github.com/tedawf/bulbsocial/internal/db"
)

type PostService struct {
	store db.Store
}

func NewPostService(store db.Store) *PostService {
	return &PostService{store: store}
}

func (p *PostService) CreatePost(ctx context.Context, params db.CreatePostParams) (db.CreatePostRow, error) {
	return p.store.CreatePost(ctx, params)
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
