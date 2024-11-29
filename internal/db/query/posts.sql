-- name: CreatePost :one
INSERT INTO
    posts (content, title, user_id, tags)
VALUES
    ($1, $2, $3, $4)
RETURNING
    *;

-- name: GetPostByID :one
SELECT
    id,
    user_id,
    title,
    content,
    created_at,
    updated_at,
    tags,
    "version"
FROM
    posts
WHERE
    id = $1;

-- name: UpdatePost :one
UPDATE posts
SET
    title = $1,
    content = $2,
    "version" = "version" + 1
WHERE
    id = $3
    AND "version" = $4
RETURNING
    "version";

-- name: DeletePost :exec
DELETE FROM posts
WHERE
    id = $1;

-- name: GetUserFeed :many
SELECT
    p.id,
    p.user_id,
    p.title,
    p."content",
    p.created_at,
    p."version",
    p.tags,
    u.username,
    COUNT(c.id) AS comments_count
FROM
    posts p
    JOIN comments c ON c.post_id = p.id
    JOIN users u ON u.id = p.user_id
    JOIN followers f ON f.user_id = p.user_id
WHERE
    f.follower_id = $1
    OR p.user_id = $1
    AND (
        p.title ILIKE '%' || $4 || '%'
        OR p.content ILIKE '%' || $4 || '%'
    )
    AND (
        p.tags @> $5
        OR $5 IS NULL
    )
GROUP BY
    p.id,
    u.username
ORDER BY
    p.created_at DESC
LIMIT
    $2
OFFSET
    $3;