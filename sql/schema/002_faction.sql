-- +goose Up
CREATE TABLE faction (
    id NUMERIC(8, 1) UNIQUE,
    name TEXT,
    faction_id TEXT,
    FOREIGN KEY (id)
    REFERENCES models(datasheet_id)
);

-- +goose Down
DROP TABLE faction;