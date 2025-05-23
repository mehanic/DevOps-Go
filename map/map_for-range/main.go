package main

import "fmt"

func main() {

	map1 := map[int]string{
		0: "first",
		1: "second",
		2: "third",
		3: "fourth",
	}

	// map is unordered, the result printed to console may be in different order
	for key, val := range map1 {
		fmt.Println("key", key, "value", val)
	}
}
