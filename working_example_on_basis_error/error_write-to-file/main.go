package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf) // set output destination of package log
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened:", err) // will write message to log.txt
	}
}
