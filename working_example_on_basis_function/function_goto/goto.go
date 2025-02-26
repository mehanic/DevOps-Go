package main

import "fmt"

func main() {
	var i int = 10
LOOP:
	for i < 30 {
		if i%3 == 0 {
			i += 2
			goto LOOP
		}
		fmt.Printf("Numbers not divisible by 3 is %d\n", i)
		i++
	}
}

// Numbers not divisible by 3 is 10
// Numbers not divisible by 3 is 11
// Numbers not divisible by 3 is 14
// Numbers not divisible by 3 is 17
// Numbers not divisible by 3 is 20
// Numbers not divisible by 3 is 23
// Numbers not divisible by 3 is 26
// Numbers not divisible by 3 is 29
