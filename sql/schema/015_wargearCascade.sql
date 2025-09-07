-- +goose Up
BEGIN;

ALTER TABLE wargear
  DROP CONSTRAINT IF EXISTS wargear_datasheet_id_fkey,
  DROP CONSTRAINT IF EXISTS wargear_models_datasheet_fk;

ALTER TABLE wargear
  ADD CONSTRAINT wargear_models_datasheet_fk
  FOREIGN KEY (datasheet_id)
  REFERENCES models (datasheet_id)
  ON DELETE CASCADE;

COMMIT;

-- +goose Down
BEGIN;

ALTER TABLE wargear
  DROP CONSTRAINT IF EXISTS wargear_models_datasheet_fk;

ALTER TABLE wargear
  ADD CONSTRAINT wargear_datasheet_id_fkey
  FOREIGN KEY (datasheet_id)
  REFERENCES models (datasheet_id)
  ON DELETE RESTRICT;

DROP INDEX IF EXISTS idx_wargear_datasheet_id;

COMMIT;