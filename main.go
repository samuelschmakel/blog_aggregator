package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	config "github.com/samuelschmakel/blog_aggregator/internal/config"
	"github.com/samuelschmakel/blog_aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Database handling
	dbConn, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	dbQueries := database.New(dbConn)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		cmds: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	// Check if arguments are provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		return
	}

	// Splitting command line arguments
	var cmd command
	cmd.Name = os.Args[1]
	cmd.Args = os.Args[2:]

	err = cmds.Run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't get user: %v", err)
		}
		return handler(s, cmd, user)
	}
}
