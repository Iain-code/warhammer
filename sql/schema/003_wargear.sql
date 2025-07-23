-- +goose Up
CREATE TABLE wargear (
    datasheet_id INT NOT NULL,
    id INT PRIMARY KEY,
    Name TEXT NOT NULL,
    Range TEXT NOT NULL,
    Type TEXT NOT NULL,
    A TEXT NOT NULL,
    BS_WS TEXT NOT NULL,
    Strength TEXT NOT NULL,
    AP INT,
    Damage TEXT NOT NULL,
    FOREIGN KEY (datasheet_id)
    REFERENCES models(datasheet_id)
);

-- +goose Down
DROP TABLE wargear;