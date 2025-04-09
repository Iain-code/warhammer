-- name: CreateRefreshToken :exec
INSERT INTO refresh_tokens (token, user_id, expires_at)
VALUES(
    $1,
    $2,
    $3
);

-- name: GetUserFromToken :one
SELECT users.* FROM users
JOIN refresh_tokens ON users.id = refresh_tokens.user_id
WHERE refresh_tokens.token = $1;


-- name: RevokeRefreshToken :exec
DELETE FROM refresh_tokens
WHERE refresh_tokens.token = $1;

