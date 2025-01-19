package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Println("Start application...")
	c := make(chan os.Signal)
	signal.Notify(c)
	log.Println("Exit with signal:", <-c)
}

// 2024/02/25 00:04:07 Start application...
// ^C2024/02/25 00:04:13 Exit with signal: interrupt
