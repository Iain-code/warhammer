-- +goose Up
CREATE TABLE roster (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    army_list JSONB NOT NULL DEFAULT '{}'::jsonb,
    enhancements text[] NOT NULL,
    name TEXT NOT NULL,
    faction TEXT NOT NULL
);

-- +goose Down
DROP TABLE roster;