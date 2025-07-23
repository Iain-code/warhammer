-- name: GetPointsForID :many
SELECT * FROM points
WHERE datasheet_id = ANY($1);