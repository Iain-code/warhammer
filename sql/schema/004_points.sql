-- +goose Up
CREATE TABLE points (
    id INT PRIMARY KEY NOT NULL,
    datasheet_id INT NOT NULL,
    line INT NOT NULL,
    description TEXT NOT NULL,
    cost INT NOT NULL,
    FOREIGN KEY (datasheet_id)
    REFERENCES models(datasheet_id)
);

-- +goose Down
DROP TABLE points;