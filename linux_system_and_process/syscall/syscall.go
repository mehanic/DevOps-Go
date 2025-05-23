// client.go
package main

import (
	"fmt"
	"net"
)

func main() {
	socketPath := "/tmp/mysocket.sock"
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	data := []byte("Hello from client")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

