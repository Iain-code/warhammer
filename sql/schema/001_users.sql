-- +goose up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    email TEXT UNIQUE NOT NULL,
    hashed_password TEXT
);

-- +goose down
DROP TABLE users;
