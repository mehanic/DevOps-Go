package main

import (
    "fmt"
)

func main() {
    message := "It was a bright cold day in April, and the clocks were striking thirteen."
    count := make(map[rune]int) // Create a map to store character counts

    // Loop over each character (rune) in the message
    for _, character := range message {
        // Increment the count for the character
        count[character]++
    }

    // Print the character counts
    fmt.Println(count)
}

//map[32:13 44:1 46:1 65:1 73:1 97:4 98:1 99:3 100:3 101:5 103:2 104:3 105:6 107:2 108:3 110:4 111:2 112:1 114:5 115:3 116:6 119:2 121:1]

// 32:13 — символ с кодом Unicode 32 (пробел) встречается 13 раз.

// 44:1 — запятая (,) встречается 1 раз.

// 46:1 — точка (.) встречается 1 раз.

// 65:1 — буква A встречается 1 раз.

// 97:4 — буква a встречается 4 раза.