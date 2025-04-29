package main

import (
	"fmt"
	"unsafe"
)

func main() {

//Commented out code: Initializes an empty string and attempts to call dump on it (currently commented out).
//Initializes the string hello and calls dump on it.
//Calls dump on string literals "hello" and "hello!".
//Iterates over each character of hello, slicing one character at a time, and calls dump on each slice.
//Calls dump on the result of converting hello to a byte slice and then back to a string.
//Calls dump on the result of converting hello to a rune slice and then back to a string.

	// empty := ""
	// dump(empty)

	hello := "hello"
	dump(hello)
	dump("hello")
	dump("hello!")

	for i := range hello {
		dump(hello[i : i+1])
	}

	dump(string([]byte(hello)))
	dump(string([]byte(hello)))
	dump(string([]rune(hello)))
}
//Defines a struct StringHeader that mimics the internal representation of a string.
//pointer: A uintptr that points to the backing array’s item (starting address of the string data).
//length: An integer representing the length of the string.

// StringHeader is used by a string value
// In practice, you should use: reflect.Header
type StringHeader struct {
	// points to a backing array's item
	pointer uintptr // where it starts
	length  int     // where it ends
}


//The dump function takes a string s as input.
//Uses the unsafe.Pointer to convert the string to a StringHeader struct.
//Prints the string along with its StringHeader representation using fmt.Printf.
// dump prints the string header of a string value
func dump(s string) {
	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v\n", s, ptr)
}


// "hello": {pointer:4915656 length:5}
// "hello": {pointer:4915656 length:5}
// "hello!": {pointer:4915911 length:6}
// "h": {pointer:4915656 length:1}
// "e": {pointer:4915657 length:1}
// "l": {pointer:4915658 length:1}
// "l": {pointer:4915659 length:1}
// "o": {pointer:4915660 length:1}
// "hello": {pointer:824633795176 length:5}
// "hello": {pointer:824633795216 length:5}
// "hello": {pointer:824633795256 length:5}
