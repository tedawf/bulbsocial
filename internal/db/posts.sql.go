// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: posts.sql

package db

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :one
INSERT INTO
    posts (user_id, title, content)
VALUES
    ($1, $2, $3)
RETURNING
    id,
    user_id,
    title,
    content,
    created_at,
    updated_at
`

type CreatePostParams struct {
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost, arg.UserID, arg.Title, arg.Content)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE
    id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getAllPosts = `-- name: GetAllPosts :many
SELECT
    id,
    user_id,
    title,
    content,
    created_at,
    updated_at
FROM
    posts
ORDER BY
    created_at DESC
LIMIT
    $1
OFFSET
    $2
`

type GetAllPostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAllPosts(ctx context.Context, arg GetAllPostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getAllPosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostByID = `-- name: GetPostByID :one
SELECT
    id,
    user_id,
    title,
    content,
    created_at,
    updated_at
FROM
    posts
WHERE
    id = $1
`

func (q *Queries) GetPostByID(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPostByID, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPostsByUser = `-- name: GetPostsByUser :many
SELECT
    id,
    user_id,
    title,
    content,
    created_at,
    updated_at
FROM
    posts
WHERE
    user_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type GetPostsByUserParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetPostsByUser(ctx context.Context, arg GetPostsByUserParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByUser, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchPosts = `-- name: SearchPosts :many
SELECT
    id,
    user_id,
    title,
    content,
    created_at,
    updated_at
FROM
    posts
WHERE
    title ILIKE '%' || $1 || '%'
    OR content ILIKE '%' || $1 || '%'
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3
`

type SearchPostsParams struct {
	Column1 sql.NullString `json:"column_1"`
	Limit   int32          `json:"limit"`
	Offset  int32          `json:"offset"`
}

func (q *Queries) SearchPosts(ctx context.Context, arg SearchPostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, searchPosts, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :exec
UPDATE posts
SET
    title = $2,
    content = $3,
    updated_at = NOW()
WHERE
    id = $1
`

type UpdatePostParams struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) error {
	_, err := q.db.ExecContext(ctx, updatePost, arg.ID, arg.Title, arg.Content)
	return err
}
