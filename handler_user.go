package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	database "github.com/samuelschmakel/blog_aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	name := cmd.Args[0]

	// Check if the user exists in the database
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("the user doesn't exist")
	}

	// The user exists, set them as the user in the config
	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Printf("%s has been set\n", s.cfg.CurrentUserName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("the register handler expects a single argument, the username")
	}

	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("duplicate user")
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("database error")
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("error creating user")
	}

	s.cfg.CurrentUserName = user.Name
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error setting user")
	}

	fmt.Printf("user %s was created\n", user.Name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting users")
	}
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error listing users: %v", err)
	}

	for i := 0; i < len(users); i++ {
		if users[i] == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", users[i])
		} else {
			fmt.Printf("* %v\n", users[i])
		}
	}
	return nil
}
