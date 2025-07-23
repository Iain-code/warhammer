-- name: CreateModel :exec
INSERT INTO models (old_id, datasheet_id, name, M, T, Sv, inv_sv, W, Ld, OC)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
);

-- name: GetModel :one
SELECT * FROM models
WHERE datasheet_id = $1;

-- name: GetModelsForFaction :many
SELECT models.* FROM models
JOIN faction ON models.datasheet_id = faction.id
WHERE faction.faction_id = $1;

-- GetOneModel :one
SELECT * FROM models
WHERE datasheet_id = $1;

-- name: GetWargearForModel :many
SELECT wargear.* FROM wargear
JOIN models ON wargear.datasheet_id = models.datasheet_id
WHERE wargear.datasheet_id = $1;

-- name: UpdateModel :one
UPDATE models
SET
  old_id = $2,
  name = $3,
  M = $4,
  T = $5,
  W = $6,
  Sv = $7,
  inv_sv = $8,
  Ld = $9,
  OC = $10
WHERE datasheet_id = $1
RETURNING *;

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

-- name: GetKeywordsForFaction :many
SELECT * FROM keywords
WHERE datasheet_id = ANY($1);