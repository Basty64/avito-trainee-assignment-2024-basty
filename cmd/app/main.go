package main

import (
	"avito-trainee-assignment-2024-basty/internal/di"
	"log"
	"net/http"
)

func main() {
	container := di.NewContainer()

	err := http.ListenAndServe("localhost:8080", container.HTTPRouter())
	if err != nil {
		log.Fatal(err)
	}
}
