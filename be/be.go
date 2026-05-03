package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/AkashKanteti/load-balancer/algo"
)

type backendServers struct {
	address []string
	algo    algo.Algo
}

func main() {
	addresses := []string{"localhost:9091", "localhost:9092", "localhost:9093", "localhost:9094"}

	for _, address := range addresses {
		listener, err := net.Listen("tcp", address)
		if err != nil {
			fmt.Println(err)
		}

		go handleListener(listener)
	}

	time.Sleep(1 * time.Minute)
}

func handleListener(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept new connection %v", err)
			continue
		}
		defer conn.Close()

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {
	req := make([]byte, 1024)
	_, err := conn.Read(req)
	if err != nil {
		log.Printf("failed to read from conn %v", err)
	}

	fmt.Printf("Received at backend %v", string(req))
}
