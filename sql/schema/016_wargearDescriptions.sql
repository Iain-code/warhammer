-- +goose Up
ALTER TABLE wargear ADD COLUMN description TEXT;
UPDATE wargear SET description = '' WHERE description IS NULL; 
ALTER TABLE wargear ALTER COLUMN description SET NOT NULL;

-- +goose Down
ALTER TABLE wargear ALTER COLUMN description DROP NOT NULL;
ALTER TABLE wargear DROP COLUMN description;