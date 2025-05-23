package main

import (
	"fmt"
	"unsafe"
)


// MAIN TYPES
// string
// bool
// int
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte - alias for uint8
// rune - alias for int32
// float32 float64
// complex64 complex128

var name string

const constName = "Андрей"

func main() {
	fmt.Println(name)
	fmt.Printf("Type: %T Value: %v\n", name, name)
	fmt.Println("-------------------------")
	name := "hello"
	fmt.Println(name)

	fmt.Println(constName)

	var hello string
	fmt.Println(hello + "iii")
	fmt.Println("-------------------------")

	hello = fmt.Sprintf("hello %s", constName)
	fmt.Println(hello)
	fmt.Printf("Type: %T Value: %v\n", hello, hello)

	fmt.Println(true)
	fmt.Printf("Type: %T Value: %v\n", true, true)
	fmt.Println("-------------------------")

	var num1 int64 = 15
	var num2 int = 15

	fmt.Println(num1 + int64(num2))

	var num3 uint8 = 1
	var num4 uint64 = 1

	fmt.Println(unsafe.Sizeof(num3))
	fmt.Println(unsafe.Sizeof(num4))
}


// Type: string Value:
// hello
// Андрей
// iii
// hello Андрей
// Type: string Value: hello Андрей
// true
// Type: bool Value: true
// 30
// 1
// 8
