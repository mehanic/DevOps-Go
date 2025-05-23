package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Exercise 1: Write a function to find the second largest number in an array of numbers.")
	fmt.Println("Example: max2Numbers([2, 1, 3, 4]) => 3")
	fmt.Println("-------------------------------------------------------------")
	inputarray(4)
}

func inputarray(n int) {
	array := []int{}
	for i := 0; i < n; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter array[%d] = ", i)
		str1, err := reader.ReadString('\n')
		str1 = strings.TrimSuffix(str1, "\n")
		a, err := strconv.Atoi(str1)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		array = append(array, a)
	}
	fmt.Println("The array you just entered contains the following elements:")
	fmt.Println(array)
	findmaxnumber(array)
	fmt.Println("-------------------------------------------------")
	findmax2number(array)
}

func findmaxnumber(mang1 []int) {
	maxnumber := mang1[0]
	for i := 0; i < len(mang1); i++ {
		if maxnumber < mang1[i] {
			maxnumber = mang1[i]
		}
	}
	fmt.Println("The largest element in the array is: ", maxnumber)
}

func findmax2number(mang2 []int) {
	sort.Ints(mang2)
	fmt.Println("The array sorted in ascending order is:", mang2)
	fmt.Println("==> THE SECOND LARGEST NUMBER IN THE ABOVE ARRAY IS: ", mang2[len(mang2)-2])
}
