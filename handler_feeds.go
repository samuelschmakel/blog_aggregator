package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feed, err := s.db.GetUserFromFeed(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}
	fmt.Println(feed)
	return nil
}
