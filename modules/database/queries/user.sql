-- name: FindOne :one
SELECT *
FROM users
WHERE id = @id;

-- name: FindMany :many
SELECT *
FROM users u
WHERE TRUE
  AND (CASE
           WHEN @name::TEXT = '' THEN TRUE
           ELSE LOWER(u.name) LIKE '%' || @name::TEXT || '%' END)
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

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