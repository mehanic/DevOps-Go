package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name        string `json:"-"`
	Age         int    `json:"change field"`
	notExported int    // Lower case is not visible to json
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

// [123 34 99 104 97 110 103 101 32 102 105 101 108 100 34 58 50 57 125]
// []uint8
// {"change field":29}
