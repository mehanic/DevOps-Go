package main

import (
	"fmt"
	"path"
	"time"
)

func main() {
	var (
		speed int
		now   time.Time
	)
	speed, now = 100, time.Now()
	fmt.Println(speed, now)
	println("--------------")

	var (
		speed1    = 100
		prevSpeed = 50
	)
	speed1, prevSpeed = prevSpeed, speed1
	fmt.Println(speed1, prevSpeed)
	println("--------------")

	var dir, file string
	dir, file = path.Split("css/main.css")
	fmt.Println("dir :", dir)
	fmt.Println("file:", file)
	println("--------------")

	var file1 string
	_, file1 = path.Split("css/main.css")
	fmt.Println("file:", file1)
	println("--------------")

	_, file2 := path.Split("css/main.css")
	// or this:
	// dir, file := path.Split("css/main.css")
	fmt.Println("file:", file2)

	println("--------------")

	color := "green"
	color = "blue"
	fmt.Println(color) //blue
	fmt.Println("-------------------- main1")
	main1()

}

func main1() {
	color := "green"
	// `"dark " + color` is an expression
	color = "dark " + color
	fmt.Println(color)

	println("--------------")
	n := 0.
	n = 3.14 * 2
	fmt.Println(n)
	println("--------------")
	var (
		perimeter     int
		width, height = 5, 6
	)
	perimeter = 2 * (width + height)
	fmt.Println(perimeter)

	println("--------------")
	var (
		lang    string
		version int
	)

	lang, version = "go", 2
	fmt.Println(lang, "version", version)

	println("--------------")

	var (
		planet string
		isTrue bool
		temp   float64
	)
	planet, isTrue, temp = "Mars", true, 19.5
	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	//-----

	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)

	//---

	age := 2
	fmt.Println(7.5 + float64(age))

	//--
	min := int8(127)
	max := int16(1000)

	fmt.Println(max + int16(min))

	//--

}

// 100 2024-10-01 19:49:48.240090166 +0200 CEST m=+0.000029432
// --------------
// 50 100
// --------------
// dir : css/
// file: main.css
// --------------
// file: main.css
// --------------
// file: main.css
// --------------
// blue
// -------------------- main1
// dark green
// --------------
// 6.28
// --------------
// 22
// --------------
// go version 2
// --------------
// Air is good on Mars
// It's true
// It is 19.5 degrees
// 10.5
// 9.5
// 1127
