// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    first_name, last_name
) VALUES (
    $1, $2
) RETURNING uuid
`

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, createUser, arg.FirstName, arg.LastName)
	var uuid pgtype.UUID
	err := row.Scan(&uuid)
	return uuid, err
}

const getUserByUUID = `-- name: GetUserByUUID :one
SELECT uuid, first_name, last_name, created_at, updated_at
FROM users
WHERE uuid = $1
`

func (q *Queries) GetUserByUUID(ctx context.Context, userUuid pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUUID, userUuid)
	var i User
	err := row.Scan(
		&i.Uuid,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT uuid, first_name, last_name, created_at, updated_at FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Uuid,
			&i.FirstName,
			&i.LastName,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
