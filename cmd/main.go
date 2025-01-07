package main

import (
	"log"

	"github.com/sugam12/go-api-crud/cmd/api"
)

func main() {

	apiServer := api.NewAPIServer("localhost:8080", nil)
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
