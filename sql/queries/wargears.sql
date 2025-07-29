-- name: UpdateWargear :one
UPDATE wargear
SET
  datasheet_id = $2,
  Name = $3,
  Range = $4,
  Type = $5,
  A = $6,
  BS_WS = $7,
  Strength = $8,
  AP = $9,
  Damage = $10
WHERE id = $1
RETURNING *;

-- name: GetWargearForModel :many
SELECT wargear.* FROM wargear
JOIN models ON wargear.datasheet_id = models.datasheet_id
WHERE wargear.datasheet_id = $1;

-- name: GetWargearForAll :many
SELECT * FROM wargear;