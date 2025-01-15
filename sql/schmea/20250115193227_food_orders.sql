-- +goose Up
-- +goose StatementBegin
CREATE TABLE foods_orders (

    order_id INT NOT NULL,
    customer_id INT NOT NULL DEFAULT 0 ,
    food_id INT NOT NULL,
    quantity INT NOT NULL DEFAULT 1,
    PRIMARY KEY (order_id,food_id),
    FOREIGN KEY (food_id) REFERENCES foods(id),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (customer_id) REFERENCES customers(id)
) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE foods_orders;
-- +goose StatementEnd
