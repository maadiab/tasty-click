-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_item_options (
    id SERIAL PRIMARY KEY,
    order_item_id INT REFERENCES order_items(id) ON DELETE CASCADE,
    option_id INT REFERENCES menu_item_options(id) ON DELETE CASCADE,
    price_difference DECIMAL(10, 2) DEFAULT 0.00
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_item_options;
-- +goose StatementEnd
