// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, password, salt)
VALUES ($1, $2, $3)
RETURNING id, email, password, salt
`

type CreateUserParams struct {
	Email    string
	Password string
	Salt     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.Password, arg.Salt)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Salt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, email, password, salt FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Salt,
	)
	return i, err
}
