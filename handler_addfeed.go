package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samuelschmakel/blog_aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("2 args are required: feed name and url")
	}

	feedName := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user: %v", err)
	}

	args := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), args)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %v", err)
	}

	fmt.Println(feed)
	return nil
}
