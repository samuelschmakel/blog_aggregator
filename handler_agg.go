package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samuelschmakel/blog_aggregator/internal/database"
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
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %v", err)
	}

	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %v", err)
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't collect feed: %v", err)
	}

	for _, item := range feedData.Channel.Item {
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       toNullString(item.Title),
			Url:         item.Link,
			Description: toNullString(item.Description),
			PublishedAt: toNullTimeFlexible(item.PubDate),
			FeedID:      feed.ID,
		}

		_, err := s.db.CreatePost(context.Background(), params)
		if err != nil {
			return fmt.Errorf("couldn't create post: %v", err)
		}
	}

	return nil
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func parseTimeWithLayouts(s string, layouts []string) (time.Time, bool) {
	for _, layout := range layouts {
		parsedTime, err := time.Parse(layout, s)
		if err == nil {
			return parsedTime, true // Successfully parsed
		}
	}
	return time.Time{}, false // Parsing failed for all layouts
}

// Convert a string to sql.NullTime, handling unknown layouts
func toNullTimeFlexible(s string) sql.NullTime {
	if s == "" {
		return sql.NullTime{Time: time.Time{}, Valid: false}
	}

	// List of possible layouts
	commonLayouts := []string{
		time.RFC3339,           // "2024-12-12T15:04:05Z"
		"2006-01-02",           // "YYYY-MM-DD"
		"02 Jan 2006 15:04:05", // "DD MMM YYYY HH:mm:ss"
		"01/02/2006",           // "MM/DD/YYYY"
		"02-01-2006",           // "DD-MM-YYYY"
		"15:04:05",             // "HH:mm:ss"
	}

	parsedTime, valid := parseTimeWithLayouts(s, commonLayouts)
	return sql.NullTime{Time: parsedTime, Valid: valid}
}
