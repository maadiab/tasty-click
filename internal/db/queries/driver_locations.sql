-- name: UpdateDriverLocation :exec
INSERT INTO driver_locations (id, driver_id, latitude, longitude, updated_at)
VALUES (gen_random_uuid(), $1, $2, $3, NOW())
ON CONFLICT (driver_id)
DO UPDATE SET latitude = EXCLUDED.latitude, longitude = EXCLUDED.longitude, updated_at = NOW();

-- name: GetDriverLocation :one
SELECT * FROM driver_locations WHERE driver_id = $1;

-- name: GetNearbyDrivers :many
SELECT * FROM driver_locations
WHERE earth_distance(
  ll_to_earth($1, $2),
  ll_to_earth(latitude, longitude)
) < $3;
