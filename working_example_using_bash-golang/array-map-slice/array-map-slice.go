package main

import (
	"fmt"
)

func main() {
	valueMap := map[string]string{
		"apple":  "10 dollars",
		"banana": "31.12 dollars",
	}

	valueMap["orange"] = "12.5 dollars"

	fmt.Printf("Apple costs %s\n", valueMap["apple"])

	fmt.Println("Keys in valueMap:")
	for key := range valueMap {
		fmt.Println(key)
	}

	keys := make([]string, 0, len(valueMap))
	for key := range valueMap {
		keys = append(keys, key)
	}

	fmt.Println("Keys using slice:")
	fmt.Println(keys)

	fmt.Println("All items in valueMap:")
	for key, value := range valueMap {
		fmt.Printf("%s: %s\n", key, value)
	}
}
