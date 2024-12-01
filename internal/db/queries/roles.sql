-- name: GetRoleByName :one
SELECT
    id,
    name,
    description,
    level
FROM
    roles
WHERE
    name = $1;
