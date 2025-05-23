package main

import "fmt"

func main() {
    // Assign the string with 10 replacing the formatting character to the variable 'x'
    x := fmt.Sprintf("There are %d types of people.", 10)

    // Assign the string "binary" to the variable 'binary'
    binary := "binary"

    // Assign the string "don't" to the variable 'do_not'
    doNot := "don't"

    // Assign the string with the variables binary and do_not replacing the formatting characters to 'y'
    y := fmt.Sprintf("Those who know %s and those who %s.", binary, doNot)

    // Print "There are 10 types of people."
    fmt.Println(x)

    // Print "Those who know binary and those who don't."
    fmt.Println(y)

    // Print "I said: 'There are 10 types of people.'"
    fmt.Printf("I said: %q.\n", x) // %q formats strings with quotes

    // Print "I also said: 'Those who know binary and those who don't.'."
    fmt.Printf("I also said: '%s'.\n", y)

    // Assign boolean false to 'hilarious'
    hilarious := false

    // Assign the string with an unevaluated formatting character to 'joke_evaluation'
    jokeEvaluation := "Isn't that joke so funny?! %v"

    // Print "Isn't that joke so funny? False"
    fmt.Printf(jokeEvaluation, hilarious) // %v formats the default representation of the value

    // Assign string to 'w'
    w := "This is the left side of..."

    // Assign string to 'e'
    e := "a string with a right side."

    // Print concatenation of 'w' and 'e'
    fmt.Println(w + e) // Using + with two strings concatenates them
}

// There are 10 types of people.
// Those who know binary and those who don't.
// I said: "There are 10 types of people.".
// I also said: 'Those who know binary and those who don't.'.
// Isn't that joke so funny?! falseThis is the left side of...a string with a right side.
