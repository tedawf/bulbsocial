package repository

import (
	"context"

	"github.com/tedawf/bulbsocial/internal/db/sqlc"
)

type FollowerRepository interface {
	Follow(ctx context.Context, followerID, userID int64) error
	Unfollow(ctx context.Context, followerID, userID int64) error
}

type followerRepo struct {
	store *Store
}

func (f *followerRepo) Follow(ctx context.Context, followerID int64, userID int64) error {
	return f.store.Follow(ctx, sqlc.FollowParams{
		UserID:     userID,
		FollowerID: followerID,
	})
}

func (f *followerRepo) Unfollow(ctx context.Context, followerID int64, userID int64) error {
	return f.store.Unfollow(ctx, sqlc.UnfollowParams{
		UserID:     userID,
		FollowerID: followerID,
	})
}

func NewFollowerRepository(store *Store) FollowerRepository {
	return &followerRepo{
		store: store,
	}
}
