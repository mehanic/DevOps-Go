package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path.")
		return
	}
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(b))
}
