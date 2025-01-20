package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Initialize the birthday map
	birthdays := map[string]string{
		"Alice": "Apr 1",
		"Bob":   "Dec 12",
		"Carol": "Mar 4",
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		// Prompt for a name
		fmt.Println("Enter a name (blank to quit):")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		// If the name is blank, exit the loop
		if name == "" {
			break
		}

		// Check if the name exists in the map
		if bday, ok := birthdays[name]; ok {
			fmt.Printf("%s is the birthday of %s\n", bday, name)
		} else {
			// If the name is not in the map, ask for the birthday
			fmt.Printf("I do not have birthday information for %s\n", name)
			fmt.Println("What is their birthday?")
			bday, _ := reader.ReadString('\n')
			bday = strings.TrimSpace(bday)

			// Add the new birthday to the map
			birthdays[name] = bday
			fmt.Println("Birthday database updated!")
		}
	}
}

// Enter a name (blank to quit):
// red
// I do not have birthday information for red
// What is their birthday?
// 30
// Birthday database updated!
// Enter a name (blank to quit):
// fff
// I do not have birthday information for fff
// What is their birthday?
// 11
// Birthday database updated!
// Enter a name (blank to quit):
