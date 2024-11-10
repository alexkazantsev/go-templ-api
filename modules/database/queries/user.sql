-- name: FindOne :one
SELECT *
FROM users
WHERE id = @id;

-- name: Create :one
INSERT INTO users (name, email, password)
VALUES (@name, @email, @password)
RETURNING *;

-- name: UpdateOne :one
UPDATE users
SET name  = @name,
    email = @email
WHERE id = @id
RETURNING *;

-- name: Exist :one
SELECT EXISTS(SELECT 1 FROM users WHERE id = @id);

-- name: Delete :exec
DELETE
FROM users
WHERE id = @id;