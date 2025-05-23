package main

import "fmt"

func main() {
	fmt.Println()

	for i := 0; i < 128; i++ {
		fmt.Printf("%q ", i)
	}
	fmt.Println()
	fmt.Println()

	// rune is alias of int32
	// because UTF-8 uses 1~4 bytes (8~32 bits)
	fmt.Println("rune is alias of int32")
	var xr []rune = []rune{'a', 'b', 'c'}
	for _, r := range xr {
		fmt.Printf("%q\t%v\n", r, r)
	}
	fmt.Println()
}


// '\x00' '\x01' '\x02' '\x03' '\x04' '\x05' '\x06' '\a' '\b' '\t' '\n' '\v' '\f' '\r' '\x0e' '\x0f' '\x10' '\x11' '\x12' '\x13' '\x14' '\x15' '\x16' '\x17' '\x18' '\x19' '\x1a' '\x1b' '\x1c' '\x1d' '\x1e' '\x1f' ' ' '!' '"' '#' '$' '%' '&' '\'' '(' ')' '*' '+' ',' '-' '.' '/' '0' '1' '2' '3' '4' '5' '6' '7' '8' '9' ':' ';' '<' '=' '>' '?' '@' 'A' 'B' 'C' 'D' 'E' 'F' 'G' 'H' 'I' 'J' 'K' 'L' 'M' 'N' 'O' 'P' 'Q' 'R' 'S' 'T' 'U' 'V' 'W' 'X' 'Y' 'Z' '[' '\\' ']' '^' '_' '`' 'a' 'b' 'c' 'd' 'e' 'f' 'g' 'h' 'i' 'j' 'k' 'l' 'm' 'n' 'o' 'p' 'q' 'r' 's' 't' 'u' 'v' 'w' 'x' 'y' 'z' '{' '|' '}' '~' '\x7f'

// rune is alias of int32
// 'a'	97
// 'b'	98
// 'c'	99
