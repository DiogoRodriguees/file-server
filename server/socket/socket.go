package socket

import (
	"bufio"
	"log"
	"net"
	"os"
)

func handleConnection(conn net.Conn) error {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	connectError := awaitSignIn(reader, conn)
	if connectError != nil {
		log.Fatal("Error on connect", connectError)
		return connectError
	}

	messageError := readMessages(conn, reader)
	if messageError != nil {
		log.Fatal("Error on message", messageError)
		return messageError
	}

	return nil
}

func readMessages(conn net.Conn, reader *bufio.Reader) error {
	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			log.Printf("Client disconnect: %s\n", conn.RemoteAddr())
			return err
		}

		log.Println("Receive message: " + message)
	}
}

func awaitSignIn(reader *bufio.Reader, conn net.Conn) error {
	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			log.Printf("Client disconnect: %s\n", conn.RemoteAddr())
			return err
		}

		if message == "connect\n" {
			return nil
		}
		log.Println("Receive message no connect: " + message)
	}

}

func CreateListener() net.Listener {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed on create listener... ", err)
		os.Exit(1)
	}

	return listener
}

func ListenerConnections(listener net.Listener) {
	log.Println("Awaiting connections ...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Connection failed")
			continue
		}
		go handleConnection(conn)
	}
}
