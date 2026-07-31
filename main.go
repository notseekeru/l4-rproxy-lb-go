package main

import (
	"log"
	"net"
	"fmt"
	"io"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer ln.Close()

	for {
		clientConn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(clientConn)
	}
}

func handleConnection(clientConn net.Conn) {
	defer clientConn.Close()
	serverConnection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Printf("Error connecting to server: %v", err)
		return
	}
	defer serverConnection.Close()

	go func() {
	clientCopy, err := io.Copy(serverConnection, clientConn)
	if err != nil {
		log.Printf("Error copying data from client to server: %v", err)
		return
	}

	fmt.Printf("Copied %d bytes from client to server\n", clientCopy)
	}()

	serverCopy, err := io.Copy(clientConn, serverConnection)
	if err != nil {
		log.Printf("Error copying data from server to client: %v", err)
		return
	}
	fmt.Printf("Copied %d bytes from server to client\n", serverCopy)
}
