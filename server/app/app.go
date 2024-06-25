package app

import (
	"log"
	"os"
	"server/socket"
)

func Run() {
	listener := socket.CreateListener()
	defer listener.Close()

	err := os.MkdirAll("uploads", 0755)
	if err != nil {
		log.Fatal("Error os make dir uploads: ", err)
		os.Exit(1)
	}

	socket.ListenerConnections(listener)
}
