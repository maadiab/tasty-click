// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (name,username,mobile,password,type)
VALUES ($1,$2,$3,$4,$5)
RETURNING id, name, username, mobile, password, type
`

type CreateUserParams struct {
	Name     string
	Username string
	Mobile   string
	Password string
	Type     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.Name,
		arg.Username,
		arg.Mobile,
		arg.Password,
		arg.Type,
	)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id =$1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getOneUser = `-- name: GetOneUser :one
SELECT FROM users WHERE id =$1
`

type GetOneUserRow struct {
}

func (q *Queries) GetOneUser(ctx context.Context, id int32) (GetOneUserRow, error) {
	row := q.db.QueryRowContext(ctx, getOneUser, id)
	var i GetOneUserRow
	err := row.Scan()
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, name, username, mobile, password, type FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
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
			&i.Username,
			&i.Mobile,
			&i.Password,
			&i.Type,
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

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET name =$1,username=$2,mobile=$3,type=$4 WHERE id =$5
`

type UpdateUserParams struct {
	Name     string
	Username string
	Mobile   string
	Type     string
	ID       int32
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Name,
		arg.Username,
		arg.Mobile,
		arg.Type,
		arg.ID,
	)
	return err
}