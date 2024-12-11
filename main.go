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
	cmds.cmds["login"] = handlerLogin
	cmds.cmds["register"] = handlerRegister
	cmds.cmds["reset"] = handlerReset
	cmds.cmds["users"] = handlerGetUsers
	cmds.cmds["agg"] = handlerAggregator

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

func handlerAggregator(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed from %s: %v", url, err)
	}

	fmt.Printf("Feed: %+v\n", rssFeed)
	return nil
}
