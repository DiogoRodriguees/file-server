package app

import (
	"bufio"
	"client/socket"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func Run() {
	conn := socket.ConnectOnServer()
	defer conn.Close()
	log.Println("Connected to server at localhost:8080")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Write your message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error on read message: ", err)
			continue
		}

		if !utf8.ValidString(message) {
			log.Fatal("Error: Invalid UTF-8 string")
			continue
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Fatal("Error on send message to server")
			continue
		}

	}
}
