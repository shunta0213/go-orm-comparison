-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: Users :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (
  name, age
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
  set name = $2,
  age = $3
WHERE id = $1
RETURNING *;