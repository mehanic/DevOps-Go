package main

import (
	"fmt"
	"syscall"
)

func main() {
	var info syscall.Sysinfo_t
	if err := syscall.Sysinfo(&info); err != nil {
		panic(err)
	}
	fmt.Printf("Total RAM: %v bytes\n", info.Totalram)
	fmt.Printf("Free RAM: %v bytes\n", info.Freeram)
	fmt.Printf("Uptime: %v seconds\n", info.Uptime)
}

