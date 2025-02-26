package main

// import (
// 	"fmt"
// 	//"gettingStartedWithGo/basics"
// 	//structspractice "gettingStartedWithGo/structs-practice"
// 	// "gettingStartedWithGo/structures"
// 	// bmicalculator "gettingStartedWithGo/bmiCalculator"
// )

import (
	"fmt"
	"math"
	"time"
)

func main() {
	IntroduceYourself(Bio{"John", "Dou", "USA", 2000})

	circ, explenation := CalcCircumference(22)

	fmt.Println(circ)
	fmt.Println(explenation)

	// bmicalculator.Bmi()

	// structures.InputUserData()

	//Init()
	//CalcCircumference()

}

type Bio struct {
	Name, LastName, Country string
	BirthYear               int
}

func IntroduceYourself(bio Bio) {
	currentYear := time.Now().Year()

	fmt.Printf("Hello, my name is %v %v, I am %v y.o. and I was born in %v\n", bio.Name, bio.LastName, currentYear-bio.BirthYear, bio.Country)

}

func CalcCircumference(radius float32) (float32, string) {
	circumference := 2 * math.Pi * radius

	return circumference, fmt.Sprintf("For a radius of %v, the circle circumference is %.2f.", radius, circumference)
}

// ╰─λ go run main.go                                         0 (0.001s) < 18:51:35
// Hello, my name is John Dou, I am 24 y.o. and I was born in USA
// 138.23009
// For a radius of 22, the circle circumference is 138.23.
