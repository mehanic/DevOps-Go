package main

import (
    "fmt"
    "time"

    "github.com/inancgumus/screen"
)

// placeholder is a type alias for a slice of strings.
// It represents a single digit or symbol in the clock.
type placeholder [5]string

// digits contains the ASCII art representations of the digits 0-9.
var digits = [...]placeholder{
    {"███", "█ █", "█ █", "█ █", "███"},
    {"  █", "  █", "  █", "  █", "  █"},
    {"███", "  █", "███", "█  ", "███"},
    {"███", "  █", "███", "  █", "███"},
    {"█ █", "█ █", "███", "  █", "  █"},
    {"███", "█  ", "███", "  █", "███"},
    {"███", "█  ", "███", "█ █", "███"},
    {"███", "  █", "  █", "  █", "  █"},
    {"███", "█ █", "███", "█ █", "███"},
    {"███", "█ █", "███", "  █", "███"},
}

// colon is the ASCII art representation of the colon symbol.
var colon = placeholder{"   ", " ░ ", "   ", " ░ ", "   "}

// dot is the ASCII art representation of the dot symbol (for split seconds).
var dot = placeholder{"   ", "   ", "   ", " ░ ", " ░ "}

func main() {
    screen.Clear()

    for {
        screen.MoveTopLeft()

        now := time.Now()
        hour, min, sec := now.Hour(), now.Minute(), now.Second()
        ssec := now.Nanosecond() / 1e8

        clock := [...]placeholder{
            digits[hour/10], digits[hour%10],
            colon,
            digits[min/10], digits[min%10],
            colon,
            digits[sec/10], digits[sec%10],
            dot,
            digits[ssec],
        }

        for line := range clock[0] {
            for index, digit := range clock {
                next := clock[index][line]
                if (digit == colon || digit == dot) && sec%2 == 0 {
                    next = "   "
                }
                fmt.Print(next, "  ")
            }
            fmt.Println()
        }

        const splitSecond = time.Second / 10
        time.Sleep(splitSecond)
    }
}

// ███    █       ███  ███       ███  ███       ███  
//   █    █         █    █       █      █       █    
// ███    █       ███  ███       ███  ███       ███  
// █      █         █  █           █  █         █ █  
// ███    █       ███  ███       ███  ███       ███  



