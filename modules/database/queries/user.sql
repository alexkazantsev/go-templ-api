-- name: FindOne :one
SELECT *
FROM users
WHERE id = @id;

-- name: Create :one
INSERT INTO users (name, email, password)
VALUES (@name, @email, @password)
RETURNING *;