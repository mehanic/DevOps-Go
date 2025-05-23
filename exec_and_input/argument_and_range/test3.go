package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	for index, value := range arguments {
		fmt.Println(index, value)
	}
}
