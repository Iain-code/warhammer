-- name: GetWargearDescriptions :many
SELECT * FROM wargearDescription
WHERE datasheet_id = $1;

-- name: UpdateWargearDescriptions :one
UPDATE wargearDescription
SET
  datasheet_id = $2,
  line = $3,
  name = $4,
  description = $5,
  type = $6
  WHERE id = $1
RETURNING *;

