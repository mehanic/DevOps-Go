package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name        string
	age         int
	notExported int // Lower case is not visible to json
}

func main() {
	p := Person{"clement", 29, 007}
	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}

// [123 125]
// []uint8
// {}
