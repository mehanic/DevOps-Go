package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const appName = "BMI Calculator"
const weightMsg = "Please enter your weight (kg): "
const heightMsg = "Please enter your height (m): "

func main() {
	Bmi()
}

func Bmi() {
	fmt.Printf("-----%v-----\n", appName)

	weight, _ := strconv.ParseFloat(prompt(weightMsg), 64)
	height, _ := strconv.ParseFloat(prompt(heightMsg), 64)

	fmt.Printf("Your BMI is: %.2f\n", calcBmi(weight, height))
}

func prompt(question string) string {
	fmt.Print(question)
	result, _ := reader.ReadString('\n')
	return strings.Replace(result, "\n", "", -1)
}

func calcBmi(weight, height float64) float64 {
	return weight / (height * height)
}
