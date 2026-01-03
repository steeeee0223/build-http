package main

import (
	"log"
	"net"

	"steeeee0223.http/internal/request"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("Error:", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting connection:", err)
			continue
		}

		req, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatal("Error", err)
		}

		req.Print()
	}
}
