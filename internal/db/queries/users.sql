-- name: CreateUser :one
INSERT INTO
    users (username, email, hashed_password)
VALUES
    ($1, $2, $3)
RETURNING
    id,
    username,
    email,
    hashed_password,
    created_at;

-- name: GetUserByID :one
SELECT
    id,
    username,
    email,
    hashed_password,
    created_at,
    password_changed_at
FROM
    users
WHERE
    id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET
    hashed_password = $2,
    password_changed_at = NOW()
WHERE
    id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE
    id = $1;

-- name: SearchUsers :many
SELECT
    id,
    username,
    email,
    created_at,
    password_changed_at
FROM
    users
WHERE
    username ILIKE '%' || $1 || '%'
ORDER BY
    created_at DESC
LIMIT
    $2
OFFSET
    $3;
