package main

import (
	"log"
	"time"

	"go-micro.dev/v5"
	"go-micro.dev/v5/broker"
)

func main() {
	service := micro.NewService(micro.Name("example.publisher"))
	service.Init()

	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	go func() {
		t := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-t.C:
				msg := &broker.Message{
					Header: map[string]string{"id": "1"},
					Body:   []byte("Hello from the publisher!"),
				}
				if err := broker.Publish("example.topic", msg); err != nil {
					log.Printf("Error publishing message: %v", err)
				}
			}
		}
	}()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
