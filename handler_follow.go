package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samuelschmakel/blog_aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("follow expects one argument, a url")
	}
	url := cmd.Args[0]

	// getting feed from url
	feed, err := s.db.GetFeedFromURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %v", err)
	}

	// creating new feed follow record
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %v", err)
	}

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	rows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %v", err)
	}

	for i := range rows {
		fmt.Println(rows[i].FeedName)
	}

	return nil
}
