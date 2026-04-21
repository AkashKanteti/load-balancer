package be

import (
	"fmt"
	"log"
	"net"

	"github.com/AkashKanteti/load-balancer/algo"
)

type backendServers struct {
	address []string
	algo    algo.Algo
}

func main() {
	addresses := []string{"localhost:9091", "localhost:9092", "localhost:9093", "localhost:9094"}

	for _, address := range addresses {
		listener, _ := net.Listen("tcp", address)

		go handleListener(listener)
	}
}

func handleListener(listener net.Listener) {
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

	fmt.Printf("ok %v", req)
}
