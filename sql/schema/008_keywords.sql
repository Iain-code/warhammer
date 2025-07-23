-- +goose Up
CREATE TABLE Keywords (
    id INT PRIMARY KEY,
    datasheet_id INT NOT NULL,
    keyword TEXT NOT NULL,
    FOREIGN KEY (datasheet_id) REFERENCES models(datasheet_id)
);

-- +goose Down
DROP TABLE Keywords;
