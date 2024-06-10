package app

import (
	"server/socket"
)

func Run() {
	listener := socket.CreateListener()
	defer listener.Close()
	socket.ListenerConnections(listener)
}
