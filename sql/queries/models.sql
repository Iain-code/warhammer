-- name: CreateModel :exec
INSERT INTO models (old_id, datasheet_id, name, M, T, Sv, inv_sv, W, Ld, OC)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
);