-- name: FollowUser :exec
INSERT INTO
    followers (follower_id, followee_id)
VALUES
    ($1, $2)
ON CONFLICT DO NOTHING;

-- name: UnfollowUser :exec
DELETE FROM followers
WHERE
    follower_id = $1
    AND followee_id = $2;

-- name: GetFollowers :many
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
    $3;

-- name: GetFollowees :many
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
    $3;