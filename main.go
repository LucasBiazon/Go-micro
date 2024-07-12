package main

import (
	"log"

	"go-micro.dev/v4"
)

func main() {
	service := micro.NewService(
		micro.Name("hello-world"),
		micro.Address(":8080"),
	)

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
