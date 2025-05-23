package main

import (
	"fmt"
	"time"
)

func main() {
OuterLoop:
	for {
		fmt.Println("This is the outer loop.")
		time.Sleep(1 * time.Second)

		for iteration := 1; iteration <= 3; iteration++ {
			fmt.Printf("This is inner loop %d.\n", iteration)
			if iteration == 2 {
				break OuterLoop // Break out of the outer loop.
			}
			time.Sleep(1 * time.Second)
		}
	}

	fmt.Println("This is the end of the script, thanks for playing!")
}

