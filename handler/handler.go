package handler

import (
	"io"
	"net"
	"os/exec"
)

// Handle manages to start a session for the client
// and create a pipe for user and shell communication.
func Handle(conn net.Conn) {
	// closing the connection after user logged out
	defer conn.Close()

	// creating the shell
	cmd := exec.Command("/bin/sh", "-i")
	// creating our pipeline
	rp, wp := io.Pipe()

	// binding stdin and stdout to user connection and pipe output
	cmd.Stdin = conn
	cmd.Stdout = wp

	// copy the input into pipeline
	go io.Copy(conn, rp)

	// executing the command
	cmd.Run()
}
