package tcpServer

import (
	"fmt"
	"net"
	// "os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

type Message struct {
	Sender string
	Data string
}

func Initialize (msg chan Message) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(conn)
		go handleConnection(conn, msg)
	}
}

// Handles incoming requests.
func handleConnection(conn net.Conn, msg chan Message) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	sender := conn.RemoteAddr().String()
	fmt.Println(sender)

	var data = Message{ Sender: string(sender) , Data: string(buf) }

	msg <- data
}