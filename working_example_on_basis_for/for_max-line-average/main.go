package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var numbers []float64

	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	fmt.Println(maximum(1,2,3,4))
	fmt.Println(inRange(1,100,3,4, 400, 5, 200))
	fmt.Println(average(1,100,3,4, 400, 5, 200))

	fmt.Printf("Avarage: %0.2f\n", average(numbers...))
}

func maximum (numbers ...float64) float64{
	max := math.Inf(-1)

	for _, number := range numbers{
		if number > max{
			max = number
		}
	}

	return max
}

func inRange (min float64, max float64, numbers ...float64) []float64{
	var result []float64

	for _, number := range numbers{
		if number >= min && number <= max{
			result = append(result, number)
		}
	}

	return result
}

func average(numbers ...float64) float64 {
	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	return sum/float64(len(numbers))
}


// 4
// [3 4 5]
// 101.85714285714286
// Avarage: NaN
