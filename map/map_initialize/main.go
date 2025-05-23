package main

import "fmt"

func main() {

	// method 1 to initialize a map
	var map1 = make(map[string]int)

	map1["key1"] = 22
	map1["key2"] = 24

	fmt.Println("map1", map1)

	// keyword "delete" can remove item from a map
	delete(map1, "key2")

	fmt.Println("map1", map1)

	// first return value is the value saved with key
	// second return value is bool, is true if key is exist
	v, ok := map1["key1"]
	fmt.Println("ok?:", ok, "value:", v)

	// first return value will be zero-value if key is not exist
	// second return value is bool, is false if key is not exist
	v, ok = map1["key2"]
	fmt.Println("ok?:", ok, "value:", v)

	// method 2 to initialize a map
	map2 := map[string]int{
		"key1": 22,
		"key2": 24,
	}

	fmt.Println("map2", map2)

	// reference type is nill if not initialized
	var map3 map[int]int
	fmt.Println("map3 is nil?", map3 == nil)
	// map3[1] = 2 // error

	// array is not reference type
	var arr [3]int
	fmt.Println("arr", arr)
	arr[1] = 2
	fmt.Println("arr", arr)
}

// map1 map[key1:22 key2:24]
// map1 map[key1:22]
// ok?: true value: 22
// ok?: false value: 0
// map2 map[key1:22 key2:24]
// map3 is nil? true
// arr [0 0 0]
// arr [0 2 0]
