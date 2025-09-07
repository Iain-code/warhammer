-- +goose Up
BEGIN;

ALTER TABLE keywords
  DROP CONSTRAINT IF EXISTS keywords_datasheet_id_fkey,
  DROP CONSTRAINT IF EXISTS keywords_models_datasheet_fk;

ALTER TABLE keywords
  ADD CONSTRAINT keywords_models_datasheet_fk
  FOREIGN KEY (datasheet_id)
  REFERENCES models (datasheet_id)
  ON DELETE CASCADE
;

COMMIT;

-- +goose Down
BEGIN;

ALTER TABLE keywords
  DROP CONSTRAINT IF EXISTS keywords_models_datasheet_fk;

ALTER TABLE keywords
  ADD CONSTRAINT keywords_datasheet_id_fkey
  FOREIGN KEY (datasheet_id)
  REFERENCES models (datasheet_id)
  ON DELETE RESTRICT;

DROP INDEX IF EXISTS idx_keywords_datasheet_id;

COMMIT;