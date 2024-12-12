package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("agg requries 1 argument: time between requests")
	}

	input := cmd.Args[0]
	timeBetweenRequests, err := time.ParseDuration(input)
	if err != nil {
		return fmt.Errorf("couldn't get time from input: %v", err)
	}

	fmt.Printf("Collecting feeds every %s", input)

	// starting scraping
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	//return nil
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %v", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %v", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("couldn't collect feed: %v", err)
	}

	for _, item := range feed.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}

	return nil
}
