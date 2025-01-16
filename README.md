# tasty-click
food order application for one restaurant

## Databse Tables
- users (id,name,username,mobile,email,password,type)
- customers (id,name,mobile,email,password)
- foods (id,name,price,picture,category)
- orders (id,customer_name,created_at,status)
- joining foods with orders foods_orders (order_id,food_id,quantity)  
- restaurant (name,address,phone)
- addrsses (customer_id,home_addr,work_addr,other_addr,description)

