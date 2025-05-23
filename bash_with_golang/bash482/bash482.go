package main

import (
	"fmt"
)

func main() {
	// Check how an exit status of 0 affects the logical operators:
	if true {
		fmt.Println("We get here because the first part is true!")
	} else {
		fmt.Println("We never see this because the first part is true :(")
	}

	// Check how an exit status of 1 affects the logical operators:
	if false {
		fmt.Println("Since we only continue after && with an exit status of 0, this is never printed.")
	} else {
		fmt.Println("Because we only continue after || with a return code that is not 0, we see this!")
	}
}

