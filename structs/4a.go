package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		panic(err)
	}
	name := "Max"
	n, err := file.WriteString(name)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

// ╰─λ go run 4.go                                                                                         0 (0.001s) < 21:28:12
// 8
