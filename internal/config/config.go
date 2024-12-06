package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	homeUrl, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error reading home url:", err)
		return Config{}, err
	}

	// Path to the JSON file
	jsonPath := homeUrl + "/.gatorconfig.json"

	// Open the file
	file, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return Config{}, err
	}

	defer file.Close()

	var config Config

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return Config{}, err
	}
	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username

	updatedJSON, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}

	homeUrl, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error reading home url:", err)
		return err
	}

	// Path to the JSON file
	jsonPath := homeUrl + "/.gatorconfig.json"

	err = os.WriteFile(jsonPath, updatedJSON, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}
