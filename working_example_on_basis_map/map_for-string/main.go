package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex, "whats range?", c)
	}
}

// Hex code for red is #FF000 whats range? map[green:#4bf745 red:#FF000 white:#ffffff]
// Hex code for green is #4bf745 whats range? map[green:#4bf745 red:#FF000 white:#ffffff]
// Hex code for white is #ffffff whats range? map[green:#4bf745 red:#FF000 white:#ffffff]
