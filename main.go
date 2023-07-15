package main

import (
	"log"
	"runners-postgresql/config"
	"runners-postgresql/server"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting Runners App...")
	log.Println("Initializing Configuration...")
	config := config.InitConfig("runners")
	log.Println("Initializing Database")
	dbHandler := server.InitDatabase(config)
	log.Println("Initializing HTTP server")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
