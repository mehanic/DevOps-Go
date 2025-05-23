package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Exercise 2: Given an array of strings, write a function to filter out the elements with the maximum length.")
	fmt.Println("Example: findMaxLengthElement[aba, aa, ad, c, vcd] => [aba, vcd]")
	fmt.Println("-------------------------------------------------")
	inputarray(5)
}

func inputarray(n int) {
	input := []string{}
	for i := 0; i < n; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter array[%d] = ", i)
		str1, err := reader.ReadString('\n')
		str1 = strings.TrimSuffix(str1, "\n")

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		input = append(input, str1)
	}
	fmt.Println("The array you just entered contains the following elements: ", input)
	findMaxLengthElement(input)
}

func findMaxLengthElement(arr []string) {
	maxnumber := len(arr[0])
	var newarr []string
	for i := 0; i < len(arr); i++ {
		if maxnumber <= len(arr[i]) {
			maxnumber = len(arr[i])
		}
	}
	for i := 0; i < len(arr); i++ {
		if maxnumber <= len(arr[i]) {
			newarr = append(newarr, arr[i])
		}
	}
	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("The elements in the array with the maximum length of %d are: %v", maxnumber, newarr)
	fmt.Println()
}
