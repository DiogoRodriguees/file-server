package socket

import (
	"log"
	"net"
	"os"
)

func ConnectOnServer() net.Conn {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("Error on connect with server: ", err)
		os.Exit(1)
	}
	return conn
}
