-- +goose Up
CREATE TABLE roster (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    army_list integer[] NOT NULL,
    enhancements text[] NOT NULL,
    name TEXT NOT NULL,
    faction TEXT NOT NULL
);

-- +goose Down
DROP TABLE roster;