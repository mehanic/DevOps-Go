package main

import (
	"fmt"
	"time"
)

func shout(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	go shout("Eyaaah")
	shout("Ewoooo")
}

// λ go run main.go                                                                        0 (0.001s) < 13:13:07
// Ewoooo
// Eyaaah
// Eyaaah
// Ewoooo
// Eyaaah
// Ewoooo
// Ewoooo
// Eyaaah
// Eyaaah
// Ewoooo
