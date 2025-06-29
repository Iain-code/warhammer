-- +goose up
CREATE TABLE points (
    id INT PRIMARY KEY NOT NULL,
    datasheet_id INT NOT NULL,
    line INT,
    description TEXT,
    cost INT
);

-- +goose down
DROP TABLE points;