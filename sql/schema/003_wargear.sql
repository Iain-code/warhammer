-- +goose up
CREATE TABLE wargear (
    datasheet_id INT NOT NULL,
    Field2 INT,
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

-- +goose down
DROP TABLE wargear;