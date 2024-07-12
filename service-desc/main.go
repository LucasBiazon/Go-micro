package main

import (
	"context"
	"fmt"

	"go-micro.dev/v5"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *interface{}, rsp *interface{}) error {
	fmt.Println("Service A was called")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("Service-A"),
	)
	service.Init()

	micro.RegisterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
