-- +goose Up
CREATE TABLE wargear (
    datasheet_id INT NOT NULL,
    Field2 INT NOT NULL,
    Name TEXT,
    Range TEXT,
    Type TEXT,
    A TEXT,
    BS_WS TEXT,
    S TEXT,
    AP INT,
    D TEXT,
    PRIMARY KEY (datasheet_id, Field2),
    FOREIGN KEY (datasheet_id) REFERENCES models(datasheet_id)
);

-- +goose Down
DROP TABLE wargear;