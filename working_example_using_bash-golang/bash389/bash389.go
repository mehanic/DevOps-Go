package main

import (
	"fmt"
)

func main() {
	CTR := 1

	// Until loop equivalent
	for CTR <= 9 {
		fmt.Printf("CTR var: %d\n", CTR)
		CTR++ // Increment the CTR variable by 1
	}

	fmt.Println("Finished")
}

