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


