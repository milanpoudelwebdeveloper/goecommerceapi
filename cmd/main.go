package main

import (
	"log"

	"github.com/milanpoudelwebdeveloper/goecommerceapi/cmd/api"
)

func main() {
	server := api.NewAPIServer(":5000", nil)
	if err := server.Run(); err != nil {
		log.Fatal("Something went wrong while running a server", err)
	}
}
