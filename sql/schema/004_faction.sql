-- +goose up
CREATE TABLE faction (
    id real UNIQUE,
    name TEXT,
    faction_id TEXT,
    FOREIGN KEY (id)
    REFERENCES models(datasheet_id)
);

-- +goose down
DROP TABLE faction;