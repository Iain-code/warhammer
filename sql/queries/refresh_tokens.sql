-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, user_id, expires_at)
VALUES(
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens
WHERE refresh_tokens.token = $1;

-- name: DeleteRefreshToken :exec
DELETE from refresh_tokens
WHERE refresh_tokens.token = $1;

