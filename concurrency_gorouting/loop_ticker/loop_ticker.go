package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for range ticker.C {
			log.Println("Hey!")
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
}

// 2025/01/28 20:58:05 Hey!
// 2025/01/28 20:58:06 Hey!
// 2025/01/28 20:58:07 Hey!
// 2025/01/28 20:58:08 Hey!
