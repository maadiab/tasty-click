-- +goose Up
-- +goose StatementBegin
CREATE TABLE menu_items (
  id SERIAL PRIMARY KEY,
  category_id INT REFERENCES categories(id),
  name VARCHAR(100) NOT NULL,
  description TEXT,
  price DECIMAL(10,2) NOT NULL,
  image_url TEXT,
  available BOOLEAN DEFAULT TRUE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE menu_items;
-- +goose StatementEnd
