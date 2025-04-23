package main

import "fmt"

type Names []string

func (n Names) Print() {
	fmt.Println(n)
}

func (n Names) AddName(name string) Names {
	return append(n, name)
}

func RollCall(names Names) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func main() {
	// Create a new Names instance
	names := Names{"Alice", "Bob"}

	// Print the initial names
	fmt.Println("Initial names:")
	names.Print()

	// Add a new name to the Names instance
	names = names.AddName("Charlie")

	// Print the names after adding a new one
	fmt.Println("\nNames after adding 'Charlie':")
	names.Print()

	// Use the RollCall function to print each name individually
	fmt.Println("\nRoll call:")
	RollCall(names)
}

// Initial names:
// [Alice Bob]

// Names after adding 'Charlie':
// [Alice Bob Charlie]

// Roll call:
// Alice
// Bob
// Charlie
