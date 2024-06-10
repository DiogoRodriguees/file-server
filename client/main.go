package main

import (
	"client/app"
	"log"
)

func main() {
	log.Println("Starting client")
	app.Run()
	log.Println("Stopping client")
}
