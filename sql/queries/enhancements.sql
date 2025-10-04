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
  description = $3,
  detachment = $4,
  faction_id = $5,
  name = $6
WHERE id = $1;


-- name: DeleteEnhancement :exec
DELETE FROM enhancements
WHERE id = $1;