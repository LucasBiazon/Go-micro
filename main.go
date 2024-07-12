package main

import (
	"fmt"
	"log"
	"net/http"

	"go-micro.dev/v5/web"
)

func main() {
	service := web.NewService(
		web.Name("hello-world"),
		web.Address(":8080"),
	)

	service.Init()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	service.Handle("/", http.DefaultServeMux)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
