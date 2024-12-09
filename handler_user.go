package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	name := cmd.Args[0]

	err := s.Config.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Printf("%s has been set\n", s.Config.CurrentUserName)
	return nil
}
