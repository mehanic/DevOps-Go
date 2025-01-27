package main

import "fmt"

func main() {

	// var a int8   //1 байт памяти 8 бит -128 127
	// var b int16  // -32768 327672 baita
	// var c int32  // 4 byte -2bill 2 bill
	// var d int64  // - 8byte -9pent
	// var e uint8  //0 255 1 byte
	// var f uint16 //0 65535 2 byte
	// var g uint32 // 4byte 4 bil
	// var h uint64 //18pent

	// var i byte //synonym uint8
	// var j rune //synonym int32
	// var k int  //libo int32 or int64
	// var m uint // bezznaka uint32 uint64

	// var a1 float32
	// var a2 float64

	// var complex complex64
	// var complex2 complex128

	// var name string = "dd"

	// var name2 = "adasd"

	// var (
	// 	name = "artur"
	// 	age  = 32
	// )
	var name, age = "rrr", 32
	name = "Dauren"
	var c string = fmt.Sprintf("My name is, %s. And i am %d year old", name, age)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(c)
}

// Dauren
// 32
// My name is, Dauren. And i am 32 year old
