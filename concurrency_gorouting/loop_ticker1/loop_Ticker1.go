package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C {
			log.Printf("Tick at: %v\n", t.UTC())
			// do something
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
}


// 2025/01/28 20:58:35 Tick at: 2025-01-28 19:58:35.256689163 +0000 UTC
// 2025/01/28 20:58:36 Tick at: 2025-01-28 19:58:36.25668902 +0000 UTC
// 2025/01/28 20:58:37 Tick at: 2025-01-28 19:58:37.25668966 +0000 UTC
// 2025/01/28 20:58:38 Tick at: 2025-01-28 19:58:38.256689043 +0000 UTC
