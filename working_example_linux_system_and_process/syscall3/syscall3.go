// readfile.go
package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fd, err := syscall.Open("/etc/passwd", syscall.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer syscall.Close(fd)

	var buf [1024]byte
	n, err := syscall.Read(fd, buf[:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
}

