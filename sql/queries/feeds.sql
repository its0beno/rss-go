-- name: Createfeed :one
INSERT INTO feeds
 ( id,created_at,update_at, name, url, user_id ) VALUES(
     $1, $2, $3, $4, $5,$6
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;