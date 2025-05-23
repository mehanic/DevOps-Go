package main

import "fmt"

func main() {
    fmt.Println("I will now count my chickens:")

    // add 25 to (30 divided by 6)
    fmt.Println("Hens", 25 + 30 / 6)

    // subtract 100 by ((25 times 3) mod 4)
    fmt.Println("Roosters", 100 - 25 * 3 % 4)

    fmt.Println("Now I will count the eggs:")

    // 3 plus 2 plus 1 minus 5 plus (4 mod 2) minus (1 divided by 4) plus six
    fmt.Println(3 + 2 + 1 - 5 + 4 % 2 - 1 / 4 + 6)

    fmt.Println("Is it true that 3 + 2 < 5 - 7?")

    // test if 5 < -2
    fmt.Println(3 + 2 < 5 - 7)

    // add 3 plus 2
    fmt.Println("What is 3 + 2?", 3 + 2)

    // subtract 7 from 5
    fmt.Println("What is 5 - 7?", 5 - 7)

    fmt.Println("Oh, that's why it's False.")

    fmt.Println("How about some more.")

    // test if 5 is greater than -2
    fmt.Println("Is it greater?", 5 > -2)

    // test if 5 is greater or equal to -2
    fmt.Println("Is it greater or equal?", 5 >= -2)

    // test if 5 is less than or equal to -2
    fmt.Println("Is it less or equal?", 5 <= -2)
}


// I will now count my chickens:
// Hens 30
// Roosters 97
// Now I will count the eggs:
// 7
// Is it true that 3 + 2 < 5 - 7?
// false
// What is 3 + 2? 5
// What is 5 - 7? -2
// Oh, that's why it's False.
// How about some more.
// Is it greater? true
// Is it greater or equal? true
// Is it less or equal? false
