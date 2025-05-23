package main

import (
	"fmt"
	"path"
)

func main() {
	_, b := multi()
	fmt.Println(b)
	main1()
	main2()
}

func multi() (int, int) {
	return 5, 4
}

func main1() {
	color, color2 := "red", "blue"
	color, color2 = "orange", "green"
	fmt.Println(color, color2)

	red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Println(red, blue)

	dir, _ := path.Split("secret/file.txt")

	fmt.Println(dir)

	speed := 100 // int
	force := 2.5 // float64

	speed = speed * int(force)
	fmt.Println(speed)

}

func main2() {
	speed := 100
	force := 2.5

	fmt.Printf("speed     : %T\n", speed)
	fmt.Printf("conversion: %T\n", float64(speed))
	fmt.Printf("expression: %T\n", float64(speed)*force)

	// CORRECT: int * int
	speed = int(float64(speed) * force)

	fmt.Println(speed)
	///---------------
	var apple int
	var orange int32

	apple = int(orange)
	orange = 65 // 65 is A
	color := string(orange)
	fmt.Println(color)
	fmt.Println(string([]byte{104, 105}))
	_ = apple

	//---

	a, b := 10, 5.5
	fmt.Println(float64(a) + b)

	//---

}

// 4
// orange green
// blue red
// secret/
// 200
// speed     : int
// conversion: float64
// expression: float64
// 250
// A
// hi
// 15.5
