package main

import (
	"fmt"
	
	"os"
	"strconv"
	
)


// break-up
func main() {
	if len(os.Args) < 3 {
		fmt.Println("gimme two numbers")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("wrong numbers")
		return
	}

	var (
		i   = min
		sum int
	)

	for {
		if i > max {
			break
		} else if i%2 != 0 {
			i++
			continue
		}

		fmt.Print(i)
		if i < max-1 {
			fmt.Print(" + ")
		}

		sum += i
		i++
	}
	fmt.Printf(" = %d\n", sum)
}
