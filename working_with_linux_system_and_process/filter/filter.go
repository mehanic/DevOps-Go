package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"Dono", "Kasino", "Indro", "Waluya"}
	var dataTerdapatO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("original data \t\t:", data)

    fmt.Println("filter contains letter O:", dataTerdapatO)

    fmt.Println("filter length is 5:", dataLength5)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
