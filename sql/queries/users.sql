-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserFromEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
