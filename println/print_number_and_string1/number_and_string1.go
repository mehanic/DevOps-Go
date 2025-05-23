package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := 10

	// ALTERNATIVES:
	// n = n - 1
	// n -= 1

	// BETTER:
	n--

	fmt.Println(n)
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
}

//--------

func main1() {
	var counter int

	// following "statements" are correct:

	counter++ // 1
	counter++ // 2
	counter++ // 3
	counter-- // 2
	fmt.Printf("There are %d line(s) in the file\n",
		counter)

	// the following "expressions" are incorrect:

	// counter = 5+counter--
	// counter = ++counter + counter--
}

func main2() {
	width, height := 5., 12.

	// calculates the area of a rectangle
	area := width * height
	fmt.Printf("%gx%g=%g\n", width, height, area)

	area = area - 10 // decreases area by 10
	area = area + 10 // increases area by 10
	area = area * 2  // doubles the area
	area = area / 2  // divides the area by 2
	fmt.Printf("area=%g\n", area)

	// // ASSIGNMENT OPERATIONS
	area -= 10 // decreases area by 10
	area += 10 // increases area by 10
	area *= 2  // doubles the area
	area /= 2  // divides the area by 2
	fmt.Printf("area=%g\n", area)

	// finds the remainder of area variable
	// since: area is float, this won't work:
	// area %= 7

	// this works
	area = float64(int(area) % 7)
	fmt.Printf("area=%g\n", area)
}

func main3() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a temperature in Celsius (e.g., 'go run main.go 25')")
		return
	}

	// Parse the input argument (convert string to float)
	c, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid temperature value:", err)
		return
	}

	// Convert Celsius to Fahrenheit
	f := c*1.8 + 32

	// Output the result
	fmt.Printf("%g ºC is %g ºF\n", c, f)
}


func main4() {
	fmt.Println(50 + 25)
	fmt.Println(50 - 15.5)
	fmt.Println(50 * 0.5)
	fmt.Println(50 / 0.5)
	fmt.Println(25 % 3)
	fmt.Println(-(5 + 2))
}

func main5() {
	// Below solutions are correct:
	x := 5. / 2
	// x := 5 / 2.
	// x := float64(5) / 2
	// x := 5 / float64(2)

	fmt.Println(x)
}

func main6() {
	// 10 + 5 - 5 - 10
	fmt.Println(10 + 5 - (5 - 10))

	// -10 + 0.5 - 1 + 5.5
	fmt.Println(-10 + 0.5 - (1 + 5.5))

	// 5 + 10*2 - 5
	fmt.Println(5 + 10*(2-5))

	// 0.5*2 - 1
	fmt.Println(0.5 * (2 - 1))

	// 3 + 1/2*10 + 4
	fmt.Println((3+1)/2*10 + 4)

	// 10 / 2 * 10 % 7
	fmt.Println(10 / 2 * (10 % 7))

	// 100 / 5 / 2
	// 5  / 2 = 2
	//  5 and 2 are integers, so, the fractional part drops
	// 5. / 2 = 2.5
	//  because 5. is a float, so the result becomes a float
	fmt.Println(100 / (5. / 2))
}


// 9
// There are 2 line(s) in the file
// 5x12=60
// area=60
// area=60
// area=4
// Please provide a temperature in Celsius (e.g., 'go run main.go 25')
// 75
// 34.5
// 25
// 100
// 1
// -7
// 2.5
// 20
// -16
// -25
// 0.5
// 24
// 15
// 40
