-- +goose Up
-- Create/attach a sequence for old_id
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_class WHERE relname = 'models_old_id_seq') THEN
    CREATE SEQUENCE models_old_id_seq;
  END IF;
END$$;

-- Seed to current max, so nextval() returns MAX+1
SELECT setval('models_old_id_seq', COALESCE((SELECT MAX(old_id) FROM models), 0));

ALTER TABLE models
  ALTER COLUMN old_id SET DEFAULT nextval('models_old_id_seq');

ALTER SEQUENCE models_old_id_seq OWNED BY models.old_id;

-- Create/attach a sequence for datasheet_id
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_class WHERE relname = 'models_datasheet_id_seq') THEN
    CREATE SEQUENCE models_datasheet_id_seq;
  END IF;
END$$;

SELECT setval('models_datasheet_id_seq', COALESCE((SELECT MAX(datasheet_id) FROM models), 0));

ALTER TABLE models
  ALTER COLUMN datasheet_id SET DEFAULT nextval('models_datasheet_id_seq');

ALTER SEQUENCE models_datasheet_id_seq OWNED BY models.datasheet_id;

-- +goose Down
ALTER TABLE models
  ALTER COLUMN old_id DROP DEFAULT,
  ALTER COLUMN datasheet_id DROP DEFAULT;

DROP SEQUENCE IF EXISTS models_old_id_seq;
DROP SEQUENCE IF EXISTS models_datasheet_id_seq;
