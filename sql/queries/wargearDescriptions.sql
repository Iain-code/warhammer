-- name: GetWargearDescriptions :many
SELECT * FROM wargearDescription
WHERE datasheet_id = $1;