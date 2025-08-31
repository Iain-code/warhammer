-- name: GetAbilities :many
SELECT * FROM abilities;

-- name: GetAbilitiesForModel :many
SELECT * FROM abilities
WHERE datasheet_id = $1;

-- name: GetAbility :one
SELECT * FROM abilities
WHERE datasheet_id = $1 AND line = $2;

-- name: UpdateAbilities :exec
UPDATE abilities
SET
  ability_id = $3,
  model = $4,
  name = $5,
  description = $6,
  type = $7,
  parameter = $8
  WHERE datasheet_id = $1 AND line = $2;