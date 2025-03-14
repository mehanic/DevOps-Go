package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	items := os.Args[1:]
	if len(items) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(items)

	var data []byte
	for i, s := range items {
		data = strconv.AppendInt(data, int64(i+1), 10)
		data = append(data, '.', ' ')
		data = append(data, s...)
		data = append(data, '\n')
	}

	err := os.WriteFile("sorted.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//go run main.go creating new go.mod: module printing-clock