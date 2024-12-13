package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/samuelschmakel/blog_aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) >= 2 {
		return fmt.Errorf("browse takes a single optional argument of limit only")
	} else if len(cmd.Args) == 1 {
		parsedInt, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("Invalid argument: %v", err)
		}
		limit = parsedInt
	}

	username := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't get user information: %v", err)
	}

	params := database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: int32(limit),
	}
	row, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error getting rows: %v", err)
	}

	for _, value := range row {
		fmt.Println(value)
	}
	return nil
}
