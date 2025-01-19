package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello world")
	arr := make([]byte, 8)

	for {
		n, err := r.Read(arr)
		fmt.Printf("%q", arr[:n])
		fmt.Println(n, err, arr)
		fmt.Println(arr[:n])
		if err == io.EOF {
			break
		}
	}
}

                                                                                                                                                    
// "Hello wo"8 <nil> [72 101 108 108 111 32 119 111]
// [72 101 108 108 111 32 119 111]
// "rld"3 <nil> [114 108 100 108 111 32 119 111]
// [114 108 100]
// ""0 EOF [114 108 100 108 111 32 119 111]
// []
