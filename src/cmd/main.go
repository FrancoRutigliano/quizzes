package main

import (
	"log"
	"quizClone/config"
	"quizClone/internal/server"
)

func main() {
	config, err := config.SetUpConfig()
	if err != nil {
		log.Fatal(err)
	}
	s := server.NewServer(config)
	s.Run()
}
