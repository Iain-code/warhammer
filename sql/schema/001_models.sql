-- +goose Up
CREATE TABLE models (
    old_id INT PRIMARY KEY,
    datasheet_id INT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    M TEXT NOT NULL,
    T TEXT NOT NULL,
    Sv TEXT NOT NULL,
    inv_sv TEXT NOT NULL,
    W INT NOT NULL,
    Ld TEXT NOT NULL,
    OC INT NOT NULL
);

-- +goose Down
DROP TABLE models;
