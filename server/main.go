package main

import (
	"log"
	"server/app"
)

func main() {
	log.Println("Starting server")
	app.Run()
	log.Println("Stopping server")
}
