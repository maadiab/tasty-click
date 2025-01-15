-- +goose Up
-- +goose StatementBegin
CREATE TABLE foods (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price INT NOT NULL,
    picture TEXT NOT NULL DEFAULT '',
    category TEXT NOT NULL DEFAULT ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE foods;
-- +goose StatementEnd
