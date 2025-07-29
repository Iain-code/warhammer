-- name: GetKeywordsForFaction :many
SELECT * FROM keywords
WHERE datasheet_id = ANY($1);