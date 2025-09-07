-- +goose Up
BEGIN;

ALTER TABLE faction
  DROP CONSTRAINT IF EXISTS faction_id_fkey,
  DROP CONSTRAINT IF EXISTS faction_models_datasheet_fk;

ALTER TABLE faction
  ADD CONSTRAINT faction_models_datasheet_fk
  FOREIGN KEY (id) REFERENCES models(datasheet_id)
  ON DELETE CASCADE;

COMMIT;

-- +goose Down
BEGIN;

ALTER TABLE faction
  DROP CONSTRAINT IF EXISTS faction_models_datasheet_fk;

ALTER TABLE faction
  ADD CONSTRAINT faction_models_datasheet_fk
  FOREIGN KEY (id) REFERENCES models(datasheet_id)
  ON DELETE RESTRICT;

COMMIT;