-- +goose Up
CREATE TABLE enhancements (
  id INTEGER PRIMARY KEY,
  faction_id TEXT NOT NULL,
  name TEXT NOT NULL,
  cost INTEGER NOT NULL,
  detachment TEXT NOT NULL,
  legend TEXT NOT NULL,
  description TEXT NOT NULL,
  field8 TEXT NOT NULL
);

-- +goose Down
DROP TABLE enhancements;