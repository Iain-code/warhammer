-- name: GetFaction :one
SELECT * FROM faction
WHERE faction.id = $1;
