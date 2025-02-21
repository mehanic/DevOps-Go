package main

import "fmt"

func main() {
    // Set a variable called days to a string with shortened day names
    days := "Mon Tue Wed Thu Fri Sat Sun"

    // Set a variable called months to a string with shortened month names, separated by \n (newline) characters
    months := "Jan\nFeb\nMar\nApr\nMay\nJun\nJul\nAug"

    // Print a string, along with days
    fmt.Println("Here's the days:", days)

    // Print a string, along with months
    fmt.Println("Here's the months:", months)

    // Print a long string with line breaks
    fmt.Println(`
There's something going on here.
With the three backticks.
We'll be able to type as much as we like.
`)
}

// Here's the days: Mon Tue Wed Thu Fri Sat Sun
// Here's the months: Jan
// Feb
// Mar
// Apr
// May
// Jun
// Jul
// Aug

// There's something going on here.
// With the three backticks.
// We'll be able to type as much as we like.

