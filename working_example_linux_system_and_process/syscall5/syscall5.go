package main

import (
	"fmt"
	"syscall"
)

func main() {
	filename := "/tmp/testfile.txt"
	flags := syscall.O_RDONLY // open file in read-only mode
	mode := uint32(0666)      // file permissions - rw-rw-rw-
	fd, err := syscall.Open(filename, flags, mode)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", filename, err)
		return
	}
	defer syscall.Close(fd)
	fmt.Printf("File %s opened successfully!\n", filename)
}

