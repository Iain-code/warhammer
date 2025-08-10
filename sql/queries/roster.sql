-- name: SaveToRoster :exec
INSERT INTO roster (id, user_id, army_list, enhancements, name, faction)
VALUES (
    $1, 
    $2, 
    $3, 
    $4,
    $5,
    $6
);

-- name: GetArmies :many
SELECT * FROM roster
WHERE user_id = $1;