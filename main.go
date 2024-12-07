package main

import (
	"log"
	"os"

	config "github.com/samuelschmakel/blog_aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	state := config.State{}
	state.Config = &cfg

	cmds := config.Commands{Cmds: make(map[string]func(*config.State, config.Command) error)}
	cmds.Cmds["login"] = config.HandlerLogin

	// Check if arguments are provided
	if len(os.Args) < 2 {
		log.Fatal("a command name is required")
	}

	// Splitting command line arguments
	var cmd config.Command
	cmd.Name = os.Args[1]

	if len(os.Args) >= 3 {
		cmd.Args = os.Args[2:len(os.Args)]
	}

	// Run the command, if it is a valid command
	handler, exists := cmds.Cmds[cmd.Name]
	if !exists {
		log.Fatalf("unknown command %s", cmd.Name)
	}

	// Execute the handler function
	err = handler(&state, cmd)
	if err != nil {
		log.Fatalf("error executing command '%s': %v", cmd.Name, err)
	}
}
