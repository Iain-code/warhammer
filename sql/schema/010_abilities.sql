-- +goose Up
CREATE TABLE abilities (
    datasheet_id INT NOT NULL,
    line INT NOT NULL,
    ability_id INT NOT NULL,
    model TEXT NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    type TEXT NOT NULL,
    parameter TEXT NOT NULL
);

-- +goose Down
DROP TABLE abilities;