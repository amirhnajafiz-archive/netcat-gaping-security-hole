package main

import (
	"log"
	"net"

	"github.com/amirhnajafiz/netcat-gaping-security-hole/handler"
)

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handler.Handle(conn)
	}
}
