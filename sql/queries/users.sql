-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE name = $1;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT name
FROM users;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetUserFromFeed :many
SELECT feeds.name, feeds.url, users.name
FROM feeds
LEFT JOIN users
ON users.id = feeds.user_id;