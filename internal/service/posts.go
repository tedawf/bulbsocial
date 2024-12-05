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

func (p *PostService) CreatePost(ctx context.Context, params db.CreatePostParams) (post db.CreatePostRow, err error) {
	return post, p.store.ExecTx(ctx, func(q db.Querier) error {
		post, err = q.CreatePost(ctx, params)
		return err
	})
}

func (u *PostService) GetPostByID(ctx context.Context, postID int64) (post db.Post, err error) {
	return post, u.store.ExecTx(ctx, func(q db.Querier) error {
		post, err = q.GetPostByID(ctx, postID)
		return err
	})
}

func (u *PostService) GetAllPosts(ctx context.Context, limit, offset int32) (posts []db.Post, err error) {
	return posts, u.store.ExecTx(ctx, func(q db.Querier) error {
		posts, err = q.GetAllPosts(ctx, db.GetAllPostsParams{
			Limit:  limit,
			Offset: offset,
		})
		return err
	})
}
