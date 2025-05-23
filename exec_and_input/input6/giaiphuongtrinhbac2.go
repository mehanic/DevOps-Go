package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Exercise 1: Solve a quadratic equation")
	fmt.Println("Enter three numbers a, b, c of type int to solve the quadratic equation a*x*x + b*x + c = 0")

	reader := bufio.NewReader(os.Stdin)

	// Input a (must not be 0)
	fmt.Print("Enter a = ")
	str1, err := reader.ReadString('\n')
	str1 = strings.TrimSuffix(str1, "\n")
	a, err := strconv.ParseFloat(str1, 64)
	if a == 0 {
		fmt.Println(err.Error())
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Input b
	fmt.Print("Enter b = ")
	str2, err := reader.ReadString('\n')
	str2 = strings.TrimSuffix(str2, "\n")
	b, err := strconv.ParseFloat(str2, 64)

	// Input c
	fmt.Print("Enter c = ")
	str3, err := reader.ReadString('\n')
	str3 = strings.TrimSuffix(str3, "\n")
	c, err := strconv.ParseFloat(str3, 64)

	fmt.Print("The quadratic equation is: " + str1 + "x^2 + " + str2 + "x + " + str3 + " = 0\n")

	var delta float64
	delta = (b * b) - 4*a*c

	switch {
	case delta == 0:
		fmt.Println("The equation has a double root x1 = x2 =", -b/(2*a))
	case delta < 0:
		fmt.Println("The equation has no real solution")
	case delta > 0:
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Println("The equation has two solutions:")
		fmt.Println("x1 =", x1)
		fmt.Println("x2 =", x2)
	}
}
