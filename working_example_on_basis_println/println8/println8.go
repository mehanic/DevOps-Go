package main

import "fmt"

func main() {
    // Define the variables with escape sequences
    tabbyCat := "\tI'm tabbed in."
    persianCat := "I'm split\non a line."
    backslashCat := "I'm \\ a \\ cat."

    // Define a multiline string with escape sequences
    fatCat := `
I'll do a list:
\t* Cat food
\t* Fishies
\t* Catnip\n\t* Grass
`

    // Print the variables
    fmt.Println(tabbyCat)
    fmt.Println(persianCat)
    fmt.Println(backslashCat)
    fmt.Println(fatCat)
}

// I'm tabbed in.
// I'm split
// on a line.
// I'm \ a \ cat.

// I'll do a list:
// \t* Cat food
// \t* Fishies
// \t* Catnip\n\t* Grass

