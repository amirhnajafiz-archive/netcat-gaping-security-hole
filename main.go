package main

import (
	"fmt"
	"log"
	"net"

	"github.com/amirhnajafiz/netcat-gaping-security-hole/handler"
)

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("listening on: 20080 ...\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handler.Handle(conn)
	}
}
