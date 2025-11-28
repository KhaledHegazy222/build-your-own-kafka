package main

import (
	"log"

	"github.com/codecrafters-io/kafka-starter-go/internal/manager"
	"github.com/codecrafters-io/kafka-starter-go/internal/metadata"
	"github.com/codecrafters-io/kafka-starter-go/internal/server"
)

func main() {
	metadata, err := metadata.LoadMetadata()
	if err != nil {
		log.Fatalf("config loading err: %v", err)
	}

	log.Printf("Metadata Loaded Successfully")

	_, err = manager.InitManger(metadata)
	if err != nil {
		log.Fatalf("manager initializing err: %v", err)
	}

	log.Printf("Manger Initialized Successfully")

	sv := server.NewServer(9092)

	err = sv.Start()
	if err != nil {
		log.Fatalf("Server start error : %s", err.Error())
	}
}
