package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/AkashKanteti/load-balancer/algo"
)

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen to client %v", err)
	}
	defer listener.Close()

	serverProps := algo.ServerProps{
		Addresses: []string{
			"localhost:9091",
			"localhost:9092",
			"localhost:9093",
			"localhost:9094",
		},
	}

	selectedAlgo := flag.String("algorithm", "round-robin", "Algorithm to use for load balancing")
	flag.Parse()

	var algorithm *algo.Algorithm
	switch *selectedAlgo {
	case "round-robin":
		rr := algo.NewRoundRobin(serverProps)
		algorithm = algo.NewAlgorithm(rr)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept new connection %v", err)
			continue
		}

		go handleConnection(conn, algorithm)
	}
}

func handleConnection(conn net.Conn, algorithm *algo.Algorithm) {
	defer conn.Close()
	req := make([]byte, 1024)
	_, err := conn.Read(req)
	if err != nil {
		log.Printf("failed to read from conn %v", err)
	}

	address := algorithm.Algo.NextServer()
	fmt.Printf("server address: %s\n", address)

	backendConn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("failed to connect to backend %v", err)
	}
	defer backendConn.Close()

	_, err = backendConn.Write(req)
	if err != nil {
		log.Printf("failed to write to backend %v", err)
	}

	_, err = conn.Write([]byte("ok\n"))
	if err != nil {
		log.Printf("failed to read from conn %v", err)
	}
}
