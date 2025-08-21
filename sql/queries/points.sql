-- name: GetPointsForID :many
SELECT * FROM points
WHERE datasheet_id = ANY($1);

-- name: UpdatePointsForID :one
UPDATE points
SET 
  id = $3,
  description = $4,
  cost = $5
  WHERE datasheet_id = $1 AND line = $2
  RETURNING *;
  