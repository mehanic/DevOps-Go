package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var catNames []string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter the name of cat %d (Or enter nothing to stop.):\n", len(catNames)+1)
		name, _ := reader.ReadString('\n')
		name = name[:len(name)-1] // Remove the newline character

		if name == "" {
			break
		}
		catNames = append(catNames, name) // Append name to slice
	}

	fmt.Println("The cat names are:")
	for _, name := range catNames {
		fmt.Println("  " + name)
	}
}

// Enter the name of cat 1 (Or enter nothing to stop.):
// 3
// Enter the name of cat 2 (Or enter nothing to stop.):
// red
// Enter the name of cat 3 (Or enter nothing to stop.):
// fer
// Enter the name of cat 4 (Or enter nothing to stop.):
// cat
// Enter the name of cat 5 (Or enter nothing to stop.):
// tu
// Enter the name of cat 6 (Or enter nothing to stop.):
// se
// Enter the name of cat 7 (Or enter nothing to stop.):

// The cat names are:
//   3
//   red
//   fer
//   cat
//   tu
//   se

