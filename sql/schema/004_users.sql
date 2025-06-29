-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    username TEXT UNIQUE NOT NULL,
    hashed_password TEXT,
    is_admin BOOLEAN NOT NULL DEFAULT False
);

-- +goose Down
DROP TABLE users;
