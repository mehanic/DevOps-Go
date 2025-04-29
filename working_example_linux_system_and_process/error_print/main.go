package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened:", err)
		// err happened: open no-file.txt: no such file or directory
		log.Println("err happened:", err)
		// Package log's function Println writes message with date and time to standard error
		// 2023/06/25 14:24:00 err happened: open no-file.txt: no such file or directory
		// try "go run main.go 2&> /dev/null"
	}
}


// err happened: open no-file.txt: no such file or directory
// 2025/02/14 15:14:39 err happened: open no-file.txt: no such file or directory

