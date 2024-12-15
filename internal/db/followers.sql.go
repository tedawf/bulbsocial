// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: followers.sql

package db

import (
	"context"
)

const followUser = `-- name: FollowUser :exec
INSERT INTO
    followers (follower_id, followee_id)
VALUES
    ($1, $2)
`

type FollowUserParams struct {
	FollowerID int64 `json:"follower_id"`
	FolloweeID int64 `json:"followee_id"`
}

func (q *Queries) FollowUser(ctx context.Context, arg FollowUserParams) error {
	_, err := q.db.ExecContext(ctx, followUser, arg.FollowerID, arg.FolloweeID)
	return err
}

const getFollowees = `-- name: GetFollowees :many
SELECT
    followee_id
FROM
    followers
WHERE
    follower_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type GetFolloweesParams struct {
	FollowerID int64 `json:"follower_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) GetFollowees(ctx context.Context, arg GetFolloweesParams) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getFollowees, arg.FollowerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var followee_id int64
		if err := rows.Scan(&followee_id); err != nil {
			return nil, err
		}
		items = append(items, followee_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowers = `-- name: GetFollowers :many
SELECT
    follower_id
FROM
    followers
WHERE
    followee_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type GetFollowersParams struct {
	FolloweeID int64 `json:"followee_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) GetFollowers(ctx context.Context, arg GetFollowersParams) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getFollowers, arg.FolloweeID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var follower_id int64
		if err := rows.Scan(&follower_id); err != nil {
			return nil, err
		}
		items = append(items, follower_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unfollowUser = `-- name: UnfollowUser :execrows
DELETE FROM followers
WHERE
    follower_id = $1
    AND followee_id = $2
RETURNING
    1
`

type UnfollowUserParams struct {
	FollowerID int64 `json:"follower_id"`
	FolloweeID int64 `json:"followee_id"`
}

func (q *Queries) UnfollowUser(ctx context.Context, arg UnfollowUserParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, unfollowUser, arg.FollowerID, arg.FolloweeID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
