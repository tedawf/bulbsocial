-- name: CreateComment :one
INSERT INTO
    comments (post_id, user_id, content)
VALUES
    ($1, $2, $3)
RETURNING
    id,
    post_id,
    user_id,
    content,
    created_at;

-- name: GetCommentsByPost :many
SELECT
    id,
    post_id,
    user_id,
    content,
    created_at
FROM
    comments
WHERE
    post_id = $1
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE
    id = $1;

-- name: SearchComments :many
SELECT
    id,
    post_id,
    user_id,
    content,
    created_at
FROM
    comments
WHERE
    content ILIKE '%' || $1 || '%'
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3;