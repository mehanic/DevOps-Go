package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Print the Fibonacci sequence with n elements")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the value of n: ")
	str, err := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	n, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("The Fibonacci sequence with the following elements:")
	for i := 1; i <= n; i++ {
		fmt.Print(fib(i), " ")
	}
	fmt.Print("\n")
}

func fib(i int) int {
	if i == 2 {
		return 1
	} else if i == 1 {
		return 0
	}
	return fib(i-1) + fib(i-2)
}
