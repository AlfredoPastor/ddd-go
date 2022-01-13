package main

import (
	"context"
	"log"

	"github.com/AlfredoPastor/ddd-go/purchasing/internal"
)

func main() {
	server, err := internal.InitializeServer()
	if err != nil {
		log.Fatal(err, "server building fail")
	}
	if err = server.Run(context.Background()); err != nil {
		log.Fatal(err, "server running fail")
	}
}
