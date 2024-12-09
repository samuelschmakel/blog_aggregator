package main

import (
	"fmt"
	"log"
	"os"

	config "github.com/samuelschmakel/blog_aggregator/internal/config"
)

type state struct {
	Config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		Config: &cfg,
	}

	cmds := commands{
		cmds: make(map[string]func(*state, command) error)}
	cmds.cmds["login"] = handlerLogin

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
