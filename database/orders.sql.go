// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: orders.sql

package database

import (
	"context"
)

const createOrder = `-- name: CreateOrder :exec
INSERT INTO orders (customer_id,created_at,updated_at,status) VALUES ($1,NOW(),NOW(),$2)
RETURNING id, customer_id, created_at, updated_at, status
`

type CreateOrderParams struct {
	CustomerID int32
	Status     string
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) error {
	_, err := q.db.ExecContext(ctx, createOrder, arg.CustomerID, arg.Status)
	return err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders WHERE id =$1
`

func (q *Queries) DeleteOrder(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, id)
	return err
}

const getOneOrder = `-- name: GetOneOrder :one
SELECT id, customer_id, created_at, updated_at, status FROM orders WHERE id=$1
`

func (q *Queries) GetOneOrder(ctx context.Context, id int32) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOneOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
	)
	return i, err
}

const getOrders = `-- name: GetOrders :many
SELECT id, customer_id, created_at, updated_at, status FROM orders
`

func (q *Queries) GetOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Status,
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

const updateOrder = `-- name: UpdateOrder :exec
UPDATE orders SET updated_at = NOW(), status=$1 WHERE id =$2
`

type UpdateOrderParams struct {
	Status string
	ID     int32
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) error {
	_, err := q.db.ExecContext(ctx, updateOrder, arg.Status, arg.ID)
	return err
}
