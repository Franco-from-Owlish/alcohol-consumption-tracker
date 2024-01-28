package main

import (
	"alcohol-consumption-tracker/internal/server"
	"log"
)

func main() {
	mainServer := server.NewServer()

	err := mainServer.Server.Run()
	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
