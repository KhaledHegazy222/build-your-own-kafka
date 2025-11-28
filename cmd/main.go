package main

import (
	"log"

	"github.com/codecrafters-io/kafka-starter-go/internal/config"
	"github.com/codecrafters-io/kafka-starter-go/internal/server"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("config loading err: %v", err)
	}

	sv := server.NewServer(9092)

	err = sv.Start()
	if err != nil {
		log.Fatalf("Server start error : %s", err.Error())
	}
}
