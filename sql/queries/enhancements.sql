-- name: GetEnhancements :many
SELECT * FROM enhancements;

-- name: GetEnhancementsForFaction :many
SELECT * FROM enhancements
WHERE faction_id = $1;
