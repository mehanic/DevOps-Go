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
	fmt.Println("Exercise 3: Generate Prime Numbers")
	fmt.Println("Enter a positive integer N < 100,000 to return an array of prime numbers <= N.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter value of N: ")
	str, err := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	n, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if n < 2 {
		fmt.Println("The number " + str + " is not a prime number!")
		return
	}
	fmt.Printf("Prime numbers: ")
	printPrimeNumbers(2, n)
}

func printPrimeNumbers(num1, num2 int) {
	if num1 < 2 || num2 < 2 {
		fmt.Println("Not a prime number!")
		return
	}
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num1)
		}
		num1++
	}
	fmt.Println()
}
