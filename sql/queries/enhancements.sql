-- name: GetEnhancements :many
SELECT * FROM enhancements;

-- name: GetEnhancementsForFaction :many
SELECT * FROM enhancements
WHERE faction_id = $1;


-- name: GetEnhancementFromId :one
SELECT * FROM enhancements
WHERE id = $1;

-- name: UpdateEnhancement :exec
UPDATE enhancements
SET 
  cost = $2,
  description = $3
WHERE id = $1;
