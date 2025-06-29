-- +goose Up
CREATE TABLE wargear (
    datasheet_id INT NOT NULL,
    id INT PRIMARY KEY,
    Name TEXT NOT NULL,
    Range TEXT,
    Type TEXT,
    A TEXT,
    BS_WS TEXT,
    Strength TEXT,
    AP TEXT,
    Damage TEXT,
    FOREIGN KEY (datasheet_id)
    REFERENCES models(datasheet_id)
);

-- +goose Down
DROP TABLE wargear;