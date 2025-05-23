package main

import "fmt"

func main() {
	// break will kick me out of the loop
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}
