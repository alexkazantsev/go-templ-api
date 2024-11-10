// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package storage

import (
	"context"

	"github.com/google/uuid"
)

const create = `-- name: Create :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email, password, created_at
`

type CreateParams struct {
	Name     string
	Email    string
	Password string
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (User, error) {
	row := q.queryRow(ctx, q.createStmt, create, arg.Name, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const delete = `-- name: Delete :exec
DELETE
FROM users
WHERE id = $1
`

func (q *Queries) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteStmt, delete, id)
	return err
}

const exist = `-- name: Exist :one
SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)
`

func (q *Queries) Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	row := q.queryRow(ctx, q.existStmt, exist, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const findMany = `-- name: FindMany :many
SELECT id, name, email, password, created_at
FROM users u
WHERE TRUE
  AND (CASE
           WHEN $1::TEXT = '' THEN TRUE
           ELSE LOWER(u.name) LIKE '%' || $1::TEXT || '%' END)
LIMIT $3 OFFSET $2
`

type FindManyParams struct {
	Name   string
	Offset int32
	Limit  int32
}

func (q *Queries) FindMany(ctx context.Context, arg FindManyParams) ([]User, error) {
	rows, err := q.query(ctx, q.findManyStmt, findMany, arg.Name, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findOne = `-- name: FindOne :one
SELECT id, name, email, password, created_at
FROM users
WHERE id = $1
`

func (q *Queries) FindOne(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.queryRow(ctx, q.findOneStmt, findOne, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const updateOne = `-- name: UpdateOne :one
UPDATE users
SET name  = $1,
    email = $2
WHERE id = $3
RETURNING id, name, email, password, created_at
`

type UpdateOneParams struct {
	Name  string
	Email string
	ID    uuid.UUID
}

func (q *Queries) UpdateOne(ctx context.Context, arg UpdateOneParams) (User, error) {
	row := q.queryRow(ctx, q.updateOneStmt, updateOne, arg.Name, arg.Email, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
