package main

import (
	"log"

	"github.com/AlfredoPastor/ddd-go/inventory/internal"
)

func main() {
	server, err := internal.InitializeServer()
	if err != nil {
		log.Fatal(err, "server building fail")
	}
	if err = server.Run(); err != nil {
		log.Fatal(err, "server running fail")
	}
}
