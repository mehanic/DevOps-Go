package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		minsPerDay = 60 * 24
		weekDays   = 7
	)

	fmt.Printf("There are %d minutes in %d weeks.\n",
		minsPerDay*weekDays*2, 2)

	main1()
	fmt.Println("---------------")
	main2()
	fmt.Println("---------------")
	main3()
	fmt.Println("---------------")
	main4()
	fmt.Println("---------------")
	main5()
	fmt.Println("---------------")
	main6()
}

func main1() {
	const (
		hoursInDay   = 24
		daysInWeek   = 7
		hoursPerWeek = hoursInDay * daysInWeek
	)

	fmt.Printf("There are %d hours in %d weeks.\n",
		hoursPerWeek*5, 5)
}

func main2() {
	const (
		home   = "Milky Way Galaxy"
		length = len(home)
	)

	fmt.Printf("There are %d characters inside %q\n",
		length, home)
}

func main3() {
	const (
		pi  = 3.14159265358979323846264
		tau = pi * 2
	)

	fmt.Printf("tau = %g\n", tau)
}

func main4() {
	// Make them typeless
	const (
		width  = 25
		height = width * 2
	)

	fmt.Printf("area = %d\n", width*height)
}

func main5() {

	const later = 10
	hours, _ := time.ParseDuration("1h")
	fmt.Printf("%s later...\n", hours*later)
}

func main6() {
	const min int32 = 1

	max := 5 + min
	// above line equals to this:
	// max := int32(5) + min

	fmt.Printf("Type of max: %T\n", max)
}

// There are 20160 minutes in 2 weeks.
// There are 840 hours in 5 weeks.
// ---------------
// There are 16 characters inside "Milky Way Galaxy"
// ---------------
// tau = 6.283185307179586
// ---------------
// area = 1250
// ---------------
// 10h0m0s later...
// ---------------
// Type of max: int32
