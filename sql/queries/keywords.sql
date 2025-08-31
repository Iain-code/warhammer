-- name: GetKeywordsForFaction :many
SELECT * FROM keywords
WHERE datasheet_id = ANY($1);

-- name: GetKeywordsForModel :many
SELECT * FROM keywords
WHERE datasheet_id = $1;