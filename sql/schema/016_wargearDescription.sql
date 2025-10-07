-- +goose Up
CREATE table wargearDescription (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    datasheet_id INT NOT NULL,
    line INT NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    type TEXT NOT NULL,
    FOREIGN KEY (datasheet_id)
    REFERENCES models(datasheet_id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE wargearDescription;