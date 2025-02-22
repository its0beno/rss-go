-- name: Createfeed :one
INSERT INTO feeds
 ( id,created_at,update_at, name, url, user_id ) VALUES(
     $1, $2, $3, $4, $5,$6
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;


-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;

-- name: MarkFeedAsFetched :one
Update feeds SET last_fetched_at = NOW() ,update_at = NOW() WHERE  id =$1 RETURNING *;