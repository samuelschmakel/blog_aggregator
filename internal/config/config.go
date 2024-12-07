package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

type State struct {
	Config *Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Cmds map[string]func(*State, Command) error
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error reading home url:", err)
		return Config{}, err
	}

	// Open the file
	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var cfg Config
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return Config{}, err
	}
	return cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	s.Config.CurrentUserName = cmd.Args[0]
	fmt.Printf("%s has been set\n", s.Config.CurrentUserName)

	err := write(*s.Config)
	if err != nil {
		return err
	}
	return nil
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.Cmds[name] = f
}

func (c *Commands) run(s *State, cmd Command) error {
	key, ok := c.Cmds[cmd.Name]
	if !ok {
		return fmt.Errorf("the command %s wasn't found", cmd.Name)
	}
	err := key(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
