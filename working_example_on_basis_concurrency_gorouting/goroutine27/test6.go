package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			fmt.Println("Inside infinite loop")
			select {
			case <-done:
				fmt.Println("Ticker is stopped")
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(2000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Just stooped the ticker")
	fmt.Scanln()
}

// λ go run main.go                                                                        0 (0.001s) < 15:43:53
// https://google.com.br  is UP
// https://golang.org  is UP
// https://apple.com.br  is UP
// https://amazon.com  is UP
// https://terra.com.br  is UP
// https://uol.com.br  is UP
// https://google.com.br  is UP
// https://amazon.com  is UP
// https://golang.org  is UP
// https://apple.com.br  is UP
// https://terra.com.br  is UP
// https://uol.com.br  is UP
// https://google.com.br  is UP
// https://amazon.com  is UP
// https://golang.org  is UP
// https://terra.com.br  is UP
// https://apple.com.br  is UP
// https://uol.com.br  is UP
// ^Csignal: interrupt
