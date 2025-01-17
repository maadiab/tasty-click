-- +goose Up
-- +goose StatementBegin
CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    addr_1 TEXT NOT NULL,
    addr_2 TEXT NOT NULL DEFAULT '',
    addr_3 TEXT NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE addresses;
-- +goose StatementEnd
