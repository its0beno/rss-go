-- name: CreatefeedFollow :one
INSERT INTO feed_follows
 ( id,created_at,update_at, feed_id, user_id ) VALUES(
     $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUsersFeeds :many
SELECT * FROM feed_follows WHERE user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id = $1 AND user_id = $2 ;