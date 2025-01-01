package main

import (
	"time"

	"github.com/google/uuid"

	"github.com/its0benp/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	Name      string `json:"name"`
	URL 	string `json:"url"`
	UserID 	uuid.UUID `json:"user_id"`
}
func databaseUserToUser(DbUser database.User) User{
	return User{
		ID:DbUser.ID,
		Name: DbUser.Name,
		CreatedAt: DbUser.CreatedAt,
		UpdateAt: DbUser.UpdateAt,	
		ApiKey: DbUser.ApiKey,
	}

}
func databaseFeedToFeed( DbFeed database.Feed) Feed{
	return Feed{
		ID: DbFeed.ID,
		CreatedAt: DbFeed.CreatedAt,
		UpdateAt: DbFeed.CreatedAt,
		Name: DbFeed.Name,
		URL: DbFeed.Url,
		UserID: DbFeed.UserID,
	}
}
func databaseFeedsToFeeds( DbFeeds []database.Feed) []Feed{
	feeds :=[] Feed{}
	for _ , DbFeed := range DbFeeds{ 
		feeds = append(feeds, databaseFeedToFeed( DbFeed) )
	}
	return feeds
}
type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
	UserID 	uuid.UUID `json:"user_id"`
}
func databaseFeedFollowToFeedFollow( DbFeed database.FeedFollow) FeedFollow{
	return FeedFollow{
		ID: DbFeed.ID,
		CreatedAt: DbFeed.CreatedAt,
		UpdateAt: DbFeed.CreatedAt,
		FeedID: DbFeed.FeedID,
		UserID: DbFeed.UserID,
	}
}



func databaseFeedFollowsToFeedFollows( DbFeedFollows []database.FeedFollow) []FeedFollow {
	feed_follows :=[] FeedFollow{}
	for _ , DbFeedFollow := range DbFeedFollows{ 
		feed_follows = append(feed_follows, databaseFeedFollowToFeedFollow(DbFeedFollow) )
	}
	return feed_follows
}