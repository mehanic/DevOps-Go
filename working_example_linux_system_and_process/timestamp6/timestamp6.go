package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	time.Sleep(2 * time.Second)

	log.Printf("main, execution time %s\n", time.Since(start))
}

