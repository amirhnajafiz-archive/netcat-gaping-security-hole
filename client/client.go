package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:20080")
	if err != nil {
		log.Fatalln(err)
	}

	tmp := make([]byte, 256) // using small tmo buffer for demonstrating

	commands := []string{"ls", "pwd"}
	for _, command := range commands {
		conn.Write([]byte(command))

		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}

			break
		}

		fmt.Println(tmp[:n])
	}
}
