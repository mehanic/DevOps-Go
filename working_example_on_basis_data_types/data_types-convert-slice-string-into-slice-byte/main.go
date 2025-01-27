package main

import "fmt"

func main() {
	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var bwords [][]byte
	for _, w := range words {
		bw := []byte(w)

		fmt.Println(bw)

		bwords = append(bwords, bw)
	}

	for _, w := range bwords {
		fmt.Println(string(w))
	}
}

// [103 111 112 104 101 114]
// [112 114 111 103 114 97 109 109 101 114]
// [103 111 32 108 97 110 103 117 97 103 101]
// [103 111 32 115 116 97 110 100 97 114 100 32 108 105 98 114 97 114 121]
// gopher
// programmer
// go language
// go standard library
