package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Println("Hey!")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- struct{}{}
	log.Println("Done")
}


// 2025/01/28 20:59:04 Hey!
// 2025/01/28 20:59:05 Hey!
// 2025/01/28 20:59:06 Hey!
// 2025/01/28 20:59:07 Hey!
// 2025/01/28 20:59:08 Hey!
// 2025/01/28 20:59:08 Done
