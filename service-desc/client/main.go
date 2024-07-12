package main

import (
	"context"
	"fmt"
	"log"

	"go-micro.dev/v5"
	"go-micro.dev/v5/client"
)

func main() {
	service := micro.NewService(
		micro.Name("Service-B"),
	)
	service.Init()

	req := service.Client().NewRequest("Service-A", "Greeter.Hello",
		&map[string]interface{}{},
		client.WithContentType("application/json"))

	rsp := &map[string]interface{}{}

	if err := service.Client().Call(context.Background(), req, rsp); err != nil {
		log.Fatalf("Error calling service A: %v", err)
	}
	fmt.Println("Successfully called service A")
}
