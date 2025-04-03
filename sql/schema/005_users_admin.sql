-- +goose Up
ALTER TABLE users
ADD COLUMN is_admin BOOLEAN NOT NULL DEFAULT False;

-- +goose Down
ALTER TABLE users
DROP COLUMN is_admin;
