-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: CreateFeedFollow :one
WITH inserted AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT inserted.*, users.name AS user_name, feeds.name AS feed_name
FROM inserted
JOIN users
ON inserted.user_id = users.id
JOIN feeds
ON inserted.feed_id = feeds.id;

-- name: GetFeedFromURL :one
SELECT *
FROM feeds
WHERE feeds.url = $1;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name AS feed_name, users.name AS user
FROM feed_follows
JOIN users
ON feed_follows.user_id = users.id
JOIN feeds
ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

-- name: Unfollow :exec
DELETE FROM feed_follows
USING feeds
WHERE feed_follows.feed_id = feeds.id
AND feed_follows.user_id = $1
AND feeds.url = $2;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;