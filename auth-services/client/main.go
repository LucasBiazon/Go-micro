package main

import (
	"context"
	"fmt"
	"log"

	"go-micro.dev/v5"
	"go-micro.dev/v5/client"
	"go-micro.dev/v5/metadata"
)

type HelloRequest struct {
}

type HelloResponse struct {
}

func main() {
	service := micro.NewService(micro.Name("hello.client"))
	service.Init()

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"Token": "valid-token",
	})

	req := client.NewRequest("hello", "Greeter.Hello", &HelloRequest{})
	rsp := &HelloResponse{}

	if err := service.Client().Call(ctx, req, rsp); err != nil {
		log.Fatalf("Error calling hello service: %v", err)
	}

	fmt.Println("Successfully called hello service")
}
