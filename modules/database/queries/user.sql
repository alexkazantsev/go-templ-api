-- name: FindOne :one
SELECT *
FROM users
WHERE id = @id;