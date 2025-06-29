-- +goose Up
CREATE TABLE models (
    old_id INT,
    datasheet_id NUMERIC(6, 1) PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    M TEXT NOT NULL,
    T INT NOT NULL,
    Sv TEXT,
    inv_sv TEXT,
    W INT,
    Ld TEXT,
    OC TEXT
);

-- +goose Down
DROP TABLE models;
