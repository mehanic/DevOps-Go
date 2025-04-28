package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

// ─λ go run tickers.go                                                                     0 (0.001s) < 03:08:04
// Tick at 2024-10-04 03:08:08.896329866 +0200 CEST m=+0.500032846
// Tick at 2024-10-04 03:08:09.396329944 +0200 CEST m=+1.000032925
// Tick at 2024-10-04 03:08:09.896330451 +0200 CEST m=+1.500033430
// Ticker stopped
