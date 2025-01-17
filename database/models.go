// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"time"
)

type Address struct {
	ID          int32
	CustomerID  int32
	Addr1       string
	Addr2       string
	Addr3       string
	Description string
}

type Customer struct {
	ID       int32
	Name     string
	Mobile   string
	Email    string
	Password string
}

type Food struct {
	ID       int32
	Name     string
	Price    int32
	Picture  string
	Category string
}

type Order struct {
	ID         int32
	CustomerID int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     string
}

type OrderFood struct {
	OrderID    int32
	CustomerID int32
	FoodID     int32
	Quantity   int32
}

type User struct {
	ID       int32
	Name     string
	Username string
	Mobile   string
	Email    string
	Password string
	Type     string
}
