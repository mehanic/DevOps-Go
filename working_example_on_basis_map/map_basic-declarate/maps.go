package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	fmt.Println("map [k1]:", m["k1"])

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("len:", len(m))

	key, value := m["k3"]
	fmt.Println("key:", key, "value", value)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

// map: map[k1:7 k2:13]
// map [k1]: 7
// len: 2
// len: 1
// key: 0 value false
// map: map[bar:2 foo:1]
