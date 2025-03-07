// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: joins.sql

package database

import (
	"context"
)

const getReciept = `-- name: GetReciept :many
SELECT order_foods.order_id, order_foods.quantity,foods.name
 from order_foods
 JOIN foods
 ON order_foods.food_id = foods.id
WHERE order_foods.order_id =$1
`

type GetRecieptRow struct {
	OrderID  int32
	Quantity int32
	Name     string
}

func (q *Queries) GetReciept(ctx context.Context, orderID int32) ([]GetRecieptRow, error) {
	rows, err := q.db.QueryContext(ctx, getReciept, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRecieptRow
	for rows.Next() {
		var i GetRecieptRow
		if err := rows.Scan(&i.OrderID, &i.Quantity, &i.Name); err != nil {
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
