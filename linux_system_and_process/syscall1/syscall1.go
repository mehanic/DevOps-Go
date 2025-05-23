// server.go
package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

func main() {
	socketPath := "/tmp/mysocket.sock"
	os.Remove(socketPath)

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer l.Close()

	syscall.Chmod(socketPath, 0777)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Received %d bytes: %s\n", n, buf[:n])
}
