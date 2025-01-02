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
	timeBetweenRequest time.Duration){
	log.Printf("Starting scraping with %d workers", concurrency)
	log.Printf("Time between requests: %s", timeBetweenRequest)
	ticker:=time.NewTicker(timeBetweenRequest)

for ;; <-ticker.C {
	feeds,err:= db.GetNextFeedsToFetch( context.Background() , int32(concurrency) )
	if err != nil {
		log.Printf("Error getting feeds to scrape %v", err)
		continue
	}
	wg := &sync.WaitGroup{}

	for _,feed := range feeds {
		wg.Add(1)
		go scrapeFeed (db,wg,feed){
			err:= scrapeFeed(feed)
			if err != nil {
				log.Printf("Error scraping feed %v: %v", feed.ID, err)
			}
		}(feed)
	}
	wg.Wait()
}}


func scrapeFeed(wg $sync.WaitGroup,	db *database.Queries,feed database.Feed) {
	defer wg.Done()
	_,err:=db.MarkFeedAsFetched(context.Background(),feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched %v: %v", feed.ID, err)
	}
	rssFeed,err :=  urlTOFeedFollows()
	if err != nil {
		log.Printf("Error fetching feed %v: %v", feed.ID, err)
		return
	}
	for _,item := range rssFeed.Channel.Item{
		log.Printf("Found Post %v", item)
	}



}