package main

import "fmt"

func main() {
	var v1 int // v1 == 0
	fmt.Println(v1)

	var v2 int = 100 // v2 == 100
	fmt.Println(v2)

	v3 := 5 // var v3 int = 5
	fmt.Println(v3)

	v4 := 5
	//v4 := 5 // ошибка если 2 раза
	v4 = 10
	fmt.Println(v4)

	var v5, v6 = 1, 2
	v5, v6 = 3, 4
	v5, v6 = v6, v5

	var v7 int
	v5, v6, v7 = 1, 2, 3

	var (
		v01 = 1
		v02 = "string"
	)
	fmt.Println(v7)
	_ = v01
	_ = v02

	//Стиль
	var EarthRadiusInMeter int = 6371000
	_ = EarthRadiusInMeter

	var earth_radius_in_meter int = 6371000
	_ = earth_radius_in_meter // integer literal

	fmt.Println(
		-200, -100, 0, 50, 100, 100,
	)
	// float literal
	fmt.Println(
		-50.5, -20.5, -0., 1., 100.56, // ...
	)

	// bool constants
	fmt.Println(
		true, false,
	)

	// string literal - utf-8
	fmt.Println(
		"", // empty - prints just a space
		"hi",

		// unicode
		"nasılsın?",   // "how are you" in turkish
		"hi there 星!", // "hi there star!"
	)

	fmt.Println(42, 8500, 344433, -2323)
	fmt.Println(3.14, 6.28, -42.)
	fmt.Println(true, false)
	fmt.Println("Hi! I'm Inanc!")
	fmt.Println("Merhaba, adım İnanç!")

	fmt.Println(0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9)
	fmt.Println(0xa, 0xb, 0xc, 0xd, 0xe, 0xf)
	fmt.Println(0x11) // 17
	fmt.Println(0x19) // 25
	fmt.Println(0x32) // 50
	fmt.Println(0x64) // 100

	// integer literal
	fmt.Println(
		-200, -100, 0, 50, 100, 100,
	)
	// float literal
	fmt.Println(
		-50.5, -20.5, -0., 1., 100.56, // ...
	)

	// bool constants
	fmt.Println(
		true, false,
	)

	// string literal - utf-8
	fmt.Println(
		"", // empty - prints just a space
		"hi",

		// unicode
		"nasılsın?",   // "how are you" in turkish
		"hi there 星!", // "hi there star!"
	)
}

// 0
// 100
// 5
// 10
// 3
// -200 -100 0 50 100 100
// -50.5 -20.5 0 1 100.56
// true false
//  hi nasılsın? hi there 星!
// 42 8500 344433 -2323
// 3.14 6.28 -42
// true false
// Hi! I'm Inanc!
// Merhaba, adım İnanç!
// 0 1 2 3 4 5 6 7 8 9
// 10 11 12 13 14 15
// 17
// 25
// 50
// 100
// -200 -100 0 50 100 100
// -50.5 -20.5 0 1 100.56
// true false
//  hi nasılsın? hi there 星!
