-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?
)
RETURNING *;
--

-- name: GetUser :one
SELECT * FROM users WHERE api_key = ?;
--
