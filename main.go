package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	serverConnection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Printf("Error connecting to server: %v", err)
		return
	}
	defer serverConnection.Close()

}
