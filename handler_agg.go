package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed from %s: %v", url, err)
	}

	fmt.Printf("Feed: %+v\n", rssFeed)
	return nil
}
