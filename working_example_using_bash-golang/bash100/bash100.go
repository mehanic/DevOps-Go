package main

import (
	"fmt"
	"log"
	"time"

	expect "github.com/google/goexpect"
)

func main() {
	// Start the interactive script (replace ./interactive.sh with your script)
	cmd := "./interactive.sh"

	// Create an Expect instance with a timeout
	e, _, err := expect.Spawn(cmd, -1, expect.Verbose(true))
	if err != nil {
		log.Fatalf("Failed to start the command: %v", err)
	}

	// Expect the prompt "enter number:"
	err = e.ExpectBatch([]expect.Batcher{
		&expect.BExp{R: "enter number:"},
	}, 5*time.Second)
	if err != nil {
		log.Fatalf("Failed to match 'enter number': %v", err)
	}

	// Send the response "1"
	err = e.Send("1\n")
	if err != nil {
		log.Fatalf("Failed to send '1': %v", err)
	}

	// Expect the prompt "enter name:"
	err = e.ExpectBatch([]expect.Batcher{
		&expect.BExp{R: "enter name:"},
	}, 5*time.Second)
	if err != nil {
		log.Fatalf("Failed to match 'enter name': %v", err)
	}

	// Send the response "hello"
	err = e.Send("hello\n")
	if err != nil {
		log.Fatalf("Failed to send 'hello': %v", err)
	}

	// Expect the end of the script (EOF)
	err = e.ExpectEOF()
	if err != nil {
		log.Fatalf("Failed to expect EOF: %v", err)
	}

	// Close the Expect session
	err = e.Close()
	if err != nil {
		log.Fatalf("Failed to close the Expect session: %v", err)
	}

	fmt.Println("Interaction with the script completed successfully.")
}

