package main

import (
	"log"
	"time"
)

func track(name string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s, execution time %s\n", name, time.Since(start))
	}
}

func main() {
	defer track("main")() // do not forget about the second parentheses
	time.Sleep(2 * time.Second)
}

