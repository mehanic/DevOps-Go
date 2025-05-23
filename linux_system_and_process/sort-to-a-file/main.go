package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	items := os.Args[1:]
	if len(items) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(items)

	var data []byte
	for _, s := range items {
		data = append(data, s...)
		data = append(data, '\n')
	}

	err := os.WriteFile("sorted.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// ╰─λ go run main.go                                                                                                                                                   0 (0.003s) < 21:22:36
// Send me some items and I will sort them
