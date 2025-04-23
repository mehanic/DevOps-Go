// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main1() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

// 32°F = 0°C
// 212°F = 100°C
