package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanup(command string) {
	fmt.Printf("I was running \"%s\" when you interrupted me.\n", command)
	fmt.Println("Quitting.")
	os.Exit(1)
}

func main() {
	// Set up channel to catch OS signals
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Goroutine to listen for signals
	go func() {
		// Wait for a signal
		sig := <-signalChannel
		cleanup(sig.String())
	}()

	// Infinite loop to mimic the original behavior
	for {
		fmt.Print("hello. ")
		time.Sleep(1 * time.Second)
		fmt.Print("my ")
		time.Sleep(1 * time.Second)
		fmt.Print("name ")
		time.Sleep(1 * time.Second)
		fmt.Print("is ")
		time.Sleep(1 * time.Second)
		fmt.Println(os.Args[0]) // Get the name of the running executable
	}
}

