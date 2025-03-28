package main

import "fmt"

func main() {
	str := "hey"
	bytes := []byte{104, 101, 121}

	// same as: []byte("hey")
	fmt.Printf(`"hey" as bytes   : %d`+"\n", []byte(str))

	// same as: string([]byte{104, 101, 121})
	fmt.Printf("bytes as string  : %q\n", string(bytes))

	// runes are unicode codepoints (numbers)
	fmt.Println()
	fmt.Printf("%c                : %[1]d\n", 'h')
	fmt.Printf("%c                : %[1]d\n", 'e')
	fmt.Printf("%c                : %[1]d\n", 'y')

	// a rune literal is typeless
	// you can put it in any numeric type
	var (
		anInt   int   = 'h'
		anInt8  int8  = 'h'
		anInt16 int16 = 'h'
		anInt32 int32 = 'h'

		// rune literal's default type is: rune
		// so, you don't need to specify it.
		// aRune   rune  = 'h'
		aRune = 'h'

		// and so on...
	)

	fmt.Println()
	fmt.Printf("rune literals are typeless:\n\t%T %T %T %T %T\n",
		anInt, anInt8, anInt16, anInt32, aRune)

	fmt.Println()

	// all are the same rune

	// beginning with go 1.13 you can type: 0b0110_1000 instead
	// fmt.Printf("%q as binary: %08[1]b\n", 0b0110_1000)
	fmt.Printf("%q in decimal: %[1]d\n", 104)
	fmt.Printf("%q in binary : %08[1]b\n", 'h')
	fmt.Printf("%q in hex    : 0x%[1]x\n", 0x68)
}

// "hey" as bytes   : [104 101 121]
// bytes as string  : "hey"

// h                : 104
// e                : 101
// y                : 121

// rune literals are typeless:
// 	int int8 int16 int32 int32

// 'h' in decimal: 104
// 'h' in binary : 01101000
// 'h' in hex    : 0x68
