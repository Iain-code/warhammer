-- name: GetPointsForID :many
SELECT * FROM points
WHERE datasheet_id = ANY($1);

-- name: GetPointsForOneID :one
SELECT * FROM points
WHERE datasheet_id = $1;

-- name: UpdatePointsForID :one
UPDATE points
SET
  datasheet_id = $2,
  line = $3,
  description = $4,
  cost = $5
  WHERE id = $1
  RETURNING *;
  