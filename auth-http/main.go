package main

import (
	"fmt"
	"net/http"

	"go-micro.dev/v5/web"
)

func main() {
	// Create a new web service
	service := web.NewService(
		web.Name("http.service"),
		web.Address(":8080"),
	)

	service.Init()

	protectedHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Access to protected resource granted!")
	}

	protectedEndpoint := AuthMiddleware(protectedHandler)

	service.HandleFunc("/protected", protectedEndpoint)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
