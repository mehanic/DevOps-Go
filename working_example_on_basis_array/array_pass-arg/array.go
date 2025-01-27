package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Tutorial on array")
	reader := bufio.NewReader(os.Stdin)
	var arr [4]int64
	fmt.Println("Enter a number")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	num, _ := strconv.ParseInt(str, 10, 32)
	arr[0] = num

	fmt.Println("Enter a number")
	str2, _ := reader.ReadString('\n')
	str2 = strings.TrimSpace(str2)
	num2, _ := strconv.ParseInt(str2, 10, 32)
	arr[1] = num2

	fmt.Println("Enter a number")
	str3, _ := reader.ReadString('\n')
	str3 = strings.TrimSpace(str3)
	num3, _ := strconv.ParseInt(str3, 10, 32)
	arr[2] = num3

	fmt.Println("Array from index 0 to 2:", arr[0], arr[1], arr[2])
	fmt.Println("Complete array :", arr)
	fmt.Println("Length of array is :", len(arr))

	var newarr = [5]string{"test1", "test2", "test3"}
	fmt.Println("New array is : ", newarr)
}

// Tutorial on array
// Enter a number
// 6
// Enter a number
// 4
// Enter a number
// 7
// Array from index 0 to 2: 6 4 7
// Complete array : [6 4 7 0]
// Length of array is : 4
// New array is :  [test1 test2 test3  ]
