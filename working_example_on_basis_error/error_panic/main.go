package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	defer func() {
		fmt.Println("first")
	}()
	defer func() {
		fmt.Println("second")
	}()
	if err != nil {
		log.Panicln("err happened:", err)
		// log.Panic calls builtin panic after printing the log message
		// builtin panic stops current goroutine and prints panic message and paniced line after running defer function in reversed order
		// panic calls os.Exit(2) to exit main thread
		// panic(err)
	}

}

// 2025/02/14 15:10:16 err happened: open no-file.txt: no such file or directory
// second
// first
// panic: err happened: open no-file.txt: no such file or directory
	

// goroutine 1 [running]:
// log.Panicln({0xc000124f10?, 0xb?, 0x0?})
// 	/home/mehanic/.gvm/gos/go1.23.0/src/log/log.go:446 +0x5a
// main.main()
// 	/home/mehanic/structure/13.error/error_panic/main.go:18 +0x9b
// exit status 2

