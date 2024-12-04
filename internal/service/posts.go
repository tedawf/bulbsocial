package service

import (
	"context"

	"github.com/tedawf/bulbsocial/internal/db"
)

type PostService struct {
	store *db.Store
}

func NewPostService(store *db.Store) *PostService {
	return &PostService{store: store}
}

func (p *PostService) CreatePost(ctx context.Context, params db.CreatePostParams) (db.Post, error) {
	var post db.Post
	return post, p.store.ExecTx(ctx, func(q *db.Queries) error {
		var err error
		post, err = q.CreatePost(ctx, params)
		return err
	})
}

func (u *PostService) GetPostByID(ctx context.Context, postID int64) (db.Post, error) {
	var post db.Post
	return post, u.store.ExecTx(ctx, func(q *db.Queries) error {
		var err error
		post, err = q.GetPostByID(ctx, postID)
		return err
	})
}

func (u *PostService) GetFeed(ctx context.Context, limit, offset int32) ([]db.Post, error) {
	var posts []db.Post
	return posts, u.store.ExecTx(ctx, func(q *db.Queries) error {
		var err error
		posts, err = q.GetAllFeed(ctx, db.GetAllFeedParams{
			Limit:  limit,
			Offset: offset,
		})
		return err
	})
}
