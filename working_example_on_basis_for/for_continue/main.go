package main

import "fmt"

func main() {
	// continue will skip the rest of the current iteration and go to the next one
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
