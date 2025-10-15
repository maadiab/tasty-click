-- +goose Up
-- +goose StatementBegin
CREATE TABLE driver_locations (
    id SERIAL PRIMARY KEY,
    driver_id INT REFERENCES drivers(id) ON DELETE CASCADE,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE drivers_locations;
-- +goose StatementEnd
