-- +goose Up
BEGIN;

ALTER TABLE points
  DROP CONSTRAINT IF EXISTS points_datasheet_id_fkey,
  DROP CONSTRAINT IF EXISTS points_id_fkey;

ALTER TABLE points
  ADD CONSTRAINT points_datasheet_id_fkey
  FOREIGN KEY (datasheet_id)
  REFERENCES models (datasheet_id)
  ON DELETE CASCADE;

COMMIT;

-- +goose Down
BEGIN;

ALTER TABLE points
  DROP CONSTRAINT IF EXISTS points_datasheet_id_fkey;

ALTER TABLE points
  ADD CONSTRAINT points_datasheet_id_fkey
  FOREIGN KEY (datasheet_id)
  REFERENCES models (datasheet_id)
  ON DELETE RESTRICT;

COMMIT;