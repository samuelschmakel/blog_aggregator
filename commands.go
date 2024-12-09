package main

import "fmt"

type command struct {
	Name string
	Args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) Run(s *state, cmd command) error {
	key, ok := c.cmds[cmd.Name]
	if !ok {
		return fmt.Errorf("the command %s wasn't found", cmd.Name)
	}
	err := key(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
