package main

import (
	"fmt"
	"log"

	"go-micro.dev/v5"
	"go-micro.dev/v5/broker"
)

func main() {
	service := micro.NewService(micro.Name("example.subscriber"))
	service.Init()

	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	_, err := broker.Subscribe("example.topic", func(p broker.Event) error {
		fmt.Printf("Received message: %s\n", string(p.Message().Body))
		return nil
	})

	if err != nil {
		log.Fatalf("Error subscribing: %v", err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
