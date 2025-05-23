package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if (len(os.Args)) == 1 {
		fmt.Fprintf(os.Stderr, "Error: specify input string\n")
		os.Exit(1)
	}
	result, err := Unpack(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Unpacked string: %s\n", result)
}




const (
	symTypeUndefined = iota
	symTypeChar
	symTypeNum
	symTypeBackslash
)

const (
	symZero = 48
	symNine = symZero + 9
)

var (
	builder                                         strings.Builder
	currentSymbol, nextPrintSymbol, prevSymbol      rune
	currentPosition, prevSymbolType, repeatSymCount int
)

func processSymbol() error {
	fmt.Printf("Processing symbol: %c - %d\n", currentSymbol, currentSymbol)
	switch {
	case currentSymbol >= symZero && currentSymbol <= symNine:
		// this is num
		switch prevSymbolType {
		case symTypeUndefined:
			return fmt.Errorf("no previous symbol before number")
		case symTypeChar, symTypeNum:
			repeatSymCount = repeatSymCount*10 + int(currentSymbol) - symZero
			fmt.Printf("Current repeat time: %d\n", repeatSymCount)
			prevSymbolType = symTypeNum
		case symTypeBackslash:
			// this is not int - this is number, interpreted as char
			nextPrintSymbol = currentSymbol
			prevSymbolType = symTypeChar
		default:
			return fmt.Errorf("unknown previous symbol: %d", currentSymbol)
		}
	case currentSymbol == '\\':
		switch prevSymbolType {
		case symTypeUndefined:
			prevSymbolType = symTypeBackslash
		case symTypeNum:
			for i := 0; i < repeatSymCount; i++ {
				builder.WriteRune(nextPrintSymbol)
			}
			repeatSymCount = 0
			prevSymbolType = symTypeBackslash
		case symTypeChar:
			builder.WriteRune(prevSymbol)
			prevSymbolType = symTypeBackslash
		case symTypeBackslash:
			nextPrintSymbol = currentSymbol
			prevSymbolType = symTypeChar
		}
	default:
		// this is char
		switch prevSymbolType {
		case symTypeUndefined:
		case symTypeChar:
			builder.WriteRune(prevSymbol)
		case symTypeNum:
			for i := 0; i < repeatSymCount; i++ {
				builder.WriteRune(nextPrintSymbol)
			}
			repeatSymCount = 0
		case symTypeBackslash:
			return fmt.Errorf("backslash before char: %c at position: %d", currentSymbol, currentPosition)
		default:
			return fmt.Errorf("unknown previous symbol: %d at position: %d", currentSymbol, currentPosition)
		}
		prevSymbolType = symTypeChar
		nextPrintSymbol = currentSymbol
	}
	prevSymbol = currentSymbol
	return nil
}
// Unpack convert symbols*number to symbols
func Unpack(input string) (string, error) {
	builder.Reset()
	prevSymbolType = symTypeUndefined
	repeatSymCount = 0
	fmt.Printf("Processing string: %s \n\n", input)
	for currentPosition, currentSymbol = range input {
		if err := processSymbol(); err != nil {
			return "", err
		}
	}
	currentSymbol = 0
	processSymbol()
	return builder.String(), nil
}


// # Repeated characters
// go run main.go "x5y2"
// # Output: xxxxxyy

// # Digits as literal characters (escaped with backslash)
// go run main.go "abc\\4"
// # Output: abc4

// # Repeated digits
// go run main.go "1\\2"
// # Output: 12

// # Backslash itself
// go run main.go "abc\\\\3"
// # Output: abc\3

// # Mixed pattern
// go run main.go "h3i2\\5"
// # Output: hhhii5


// main.go "abc\\4"
// Processing string: abc\4 

// Processing symbol: a - 97
// Processing symbol: b - 98
// Processing symbol: c - 99
// Processing symbol: \ - 92
// Processing symbol: 4 - 52
// Processing symbol:  - 0
// Unpacked string: abc4


// main.go "a3b2c4"
// Processing string: a3b2c4 

// Processing symbol: a - 97
// Processing symbol: 3 - 51
// Current repeat time: 3
// Processing symbol: b - 98
// Processing symbol: 2 - 50
// Current repeat time: 2
// Processing symbol: c - 99
// Processing symbol: 4 - 52
// Current repeat time: 4
// Processing symbol:  - 0
// Unpacked string: aaabbcccc
