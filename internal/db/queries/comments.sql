-- name: CreateComment :one
INSERT INTO
    comments (post_id, user_id, content)
VALUES
    ($1, $2, $3)
RETURNING
    id,
    created_at;

-- name: GetCommentsByPostID :many
SELECT
    c.id,
    c.post_id,
    c.user_id,
    c.content,
    c.created_at,
    users.username,
    users.id
FROM
    comments c
    JOIN users ON users.id = c.user_id
WHERE
    c.post_id = $1
ORDER BY
    c.created_at DESC;
