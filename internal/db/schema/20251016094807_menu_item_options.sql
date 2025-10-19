-- +goose Up
-- +goose StatementBegin
CREATE TABLE menu_item_options (
    id SERIAL PRIMARY KEY,
    menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) DEFAULT 0
   );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE menu_item_options;
-- +goose StatementEnd
