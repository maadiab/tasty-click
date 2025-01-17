-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    username TEXT NOT NULL,
    mobile TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    type TEXT NOT NULL
) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
