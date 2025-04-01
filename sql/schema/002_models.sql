-- +goose up
CREATE TABLE models (
    datasheet_id INT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    M TEXT NOT NULL,
    T INT NOT NULL,
    Sv TEXT,
    inv_sv TEXT,
    W INT,
    Ld TEXT,
    OC TEXT
);

-- +goose down
DROP TABLE models;
