-- name: CreateUser :one
INSERT INTO
    users (username, email, password, role_id)
VALUES
    (
        $1,
        $2,
        $3,
        (
            SELECT
                id
            FROM
                roles
            WHERE
                name = $4
        )
    )
RETURNING
    *;

-- name: GetUserByID :one
SELECT
    *
FROM
    users
WHERE
    users.id = $1;

-- name: GetUserByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = $1
    AND is_verified = TRUE;

-- name: UpdateUser :one
UPDATE users
SET
    username = $1,
    email = $2,
    is_verified = $3
WHERE
    id = $4
RETURNING
    *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE
    id = $1;

-- name: GetUserFromInvitation :one
SELECT
    u.id,
    u.username,
    u.email,
    u.created_at,
    u.is_verified
FROM
    users u
    JOIN user_verifications uv ON u.id = uv.user_id
WHERE
    uv.token = $1
    AND uv.expiry > $2;
