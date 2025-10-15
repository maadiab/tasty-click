# tasty-click
food order application for one restaurant

## Databse Tables
- users (id,name,username,mobile,email,password,type)
- customers (id,name,mobile,email,password)
- foods (id,name,price,picture,category)
- orders (id,customer_id,created_at,updated_at,status)
- joining foods with orders foods_orders (order_id,customer_id,food_id,quantity)  
- setting (name,address,phone)
- addrsses (id,customer_id,home_addr,work_addr,other_addr,description)

## Technology used in this app
- Golang
- Postgresql
- Goose
- Sqlc
