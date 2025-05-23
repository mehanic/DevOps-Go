package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func setup() {
	fmt.Printf("PID of script is %d\n", os.Getpid())
}

func cleanup() {
	fmt.Println("cleaning up")
	os.Exit(1)
}

func main() {
	// Setup signal handling
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Call setup function
	setup()

	// Run cleanup on receiving signal
	go func() {
		<-sigs
		cleanup()
	}()

	// Loop forever with a noop (sleeping)
	for {
		time.Sleep(1 * time.Second)
	}
}

