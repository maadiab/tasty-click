-- +goose Up
-- +goose StatementBegin
CREATE TABLE payments (
  id SERIAL PRIMARY KEY,
  order_id INT REFERENCES orders(id) ON DELETE CASCADE,
  amount DECIMAL(10,2),
  status VARCHAR(20) DEFAULT 'pending',
  transaction_id VARCHAR(100),
  payment_provider VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payments;
-- +goose StatementEnd
