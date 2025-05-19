package main

import "fmt"

// ANSI escape codes for colors and formatting
const (
	RESET    = "\033[0m"
	RED      = "\033[1;31m"
	GREEN    = "\033[1;32m"
	YELLOW   = "\033[1;33m"
	BLUE     = "\033[1;34m"
	MAGENTA  = "\033[1;35m"
	CYAN     = "\033[1;36m"
	DARK_RED = "\033[2;31m"
	FAINT_RED = "\033[3;31m"
	UNDERLINE_RED = "\033[4;31m"
	BLACK_BG = "\033[1;40m"
	GREEN_BG = "\033[1;42m"
	YELLOW_BG = "\033[1;43m"
	BLUE_BG = "\033[1;44m"
	MAGENTA_BG = "\033[1;45m"
	CYAN_BG = "\033[1;46m"
	WHITE_BG = "\033[1;47m"
	GRAY_BG = "\033[1;48m"
)

func main() {
	// Output colored text
	fmt.Println(RED + "This is red text" + RESET)
	fmt.Println(GREEN + "This is green text" + RESET)
	fmt.Println(YELLOW + "This is yellow text" + RESET)
	fmt.Println(BLUE + "This is blue text" + RESET)
	fmt.Println(MAGENTA + "This is magenta text" + RESET)
	fmt.Println(CYAN + "This is cyan text" + RESET)
	fmt.Println(DARK_RED + "This is dark red text" + RESET)
	fmt.Println(FAINT_RED + "This is faint red text" + RESET)
	fmt.Println(UNDERLINE_RED + "This is underlined red text" + RESET)
	fmt.Println(BLACK_BG + "This is text on a black background" + RESET)
	fmt.Println(GREEN_BG + "This is text on a green background" + RESET)
	fmt.Println(YELLOW_BG + "This is text on a yellow background" + RESET)
	fmt.Println(BLUE_BG + "This is text on a blue background" + RESET)
	fmt.Println(MAGENTA_BG + "This is text on a magenta background" + RESET)
	fmt.Println(CYAN_BG + "This is text on a cyan background" + RESET)
	fmt.Println(WHITE_BG + "This is text on a white background" + RESET)
	fmt.Println(GRAY_BG + "This is text on a gray background" + RESET)
}

