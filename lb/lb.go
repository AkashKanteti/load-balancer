package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "")
	if err != nil {
		log.Fatalf("failed to listen to client %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept new connection %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	req := make([]byte, 1024)
	_, err := conn.Read(req)
	if err != nil {
		log.Printf("failed to read from conn %v", err)
	}

}
