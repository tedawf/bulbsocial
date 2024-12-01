-- name: CreateUserVerification :one
INSERT INTO
    user_verifications (token, user_id, expiry)
VALUES
    ($1, $2, $3)
RETURNING
    *;

-- name: DeleteUserVerification :exec
DELETE FROM user_verifications
WHERE
    user_id = $1;
