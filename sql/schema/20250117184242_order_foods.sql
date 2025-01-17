-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_foods (

    order_id INT NOT NULL,
    customer_id INT NOT NULL,
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
DROP TABLE order_foods;
-- +goose StatementEnd
