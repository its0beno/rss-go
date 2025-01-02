// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feed_follow.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createfeedFollow = `-- name: CreatefeedFollow :one
INSERT INTO feed_follows
 ( id,created_at,update_at, feed_id, user_id ) VALUES(
     $1, $2, $3, $4, $5
)
RETURNING id, created_at, update_at, feed_id, user_id
`

type CreatefeedFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdateAt  time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) CreatefeedFollow(ctx context.Context, arg CreatefeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createfeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdateAt,
		arg.FeedID,
		arg.UserID,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdateAt,
		&i.FeedID,
		&i.UserID,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id = $1 AND user_id = $2
`

type DeleteFeedFollowParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, arg.ID, arg.UserID)
	return err
}

const getUsersFeeds = `-- name: GetUsersFeeds :many
SELECT id, created_at, update_at, feed_id, user_id FROM feed_follows WHERE user_id = $1
`

func (q *Queries) GetUsersFeeds(ctx context.Context, userID uuid.UUID) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getUsersFeeds, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdateAt,
			&i.FeedID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
