package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/its0benp/rssagg/internal/database"
)

func startScraping(
	db *database.Queries, 
	concurrency int,
	timeBetweenRequest time.Duration) {
	log.Printf("Starting scraping with %d workers", concurrency)
	log.Printf("Time between requests: %s", timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	for range ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error getting feeds to scrape %v", err)
			continue
		}
		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go func(feed database.Feed) {
				defer wg.Done()
				err := scrapeFeed(db, feed)
				if err != nil {
					log.Printf("Error scraping feed %v: %v", feed.ID, err)
				}
			}(feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, feed database.Feed) error {
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched %v: %v", feed.ID, err)
		return err
	}
	rssFeed, err := urlTOFeedFollows("https://www.wagslane.dev/index.xml")
	if err != nil {
		log.Printf("Error fetching feed %v: %v", feed.ID, err)
		return err
	}
	for _, item := range rssFeed.Channel.Item {
		log.Printf("Found Post %v", item)
	}
	return nil
}