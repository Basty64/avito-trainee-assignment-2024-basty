package main

import (
	"avito-trainee-assignment-2024-basty/internal/app"
	"log"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}

}
