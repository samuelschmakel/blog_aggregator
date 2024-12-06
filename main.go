package main

import (
	"fmt"

	config "github.com/samuelschmakel/blog_aggregator/internal/config"
)

func main() {
	initConfig := config.Read()
	fmt.Println(initConfig)
	initConfig.SetUser("Sam")
	newConfig := config.Read()
	fmt.Println(newConfig)
}
