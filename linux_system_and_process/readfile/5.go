package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

//Information