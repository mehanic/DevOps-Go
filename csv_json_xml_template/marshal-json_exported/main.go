package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name        string // Capital is visible to json
	Age         int
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


// [123 34 78 97 109 101 34 58 34 99 108 101 109 101 110 116 34 44 34 65 103 101 34 58 50 57 125]
// []uint8
// {"Name":"clement","Age":29}
