-- +goose Up
CREATE TABLE faction (
    id INT PRIMARY KEY,
    name TEXT,
    faction_id TEXT,
    FOREIGN KEY (id)
    REFERENCES models(datasheet_id)
);

-- +goose Down
DROP TABLE faction;