-- name: CreatePost :one
INSERT INTO
    posts (user_id, title, content)
VALUES
    ($1, $2, $3)
RETURNING
    id,
    user_id,
    title,
    content,
    created_at;

-- name: GetPostByID :one
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
    id = $1;

-- name: GetPostsByUser :many
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
    $3;

-- name: UpdatePost :exec
UPDATE posts
SET
    title = $2,
    content = $3,
    updated_at = NOW()
WHERE
    id = $1;

-- name: DeletePost :exec
DELETE FROM posts
WHERE
    id = $1;

-- name: GetAllPosts :many
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
    $2;

-- name: SearchPosts :many
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
    $3;
