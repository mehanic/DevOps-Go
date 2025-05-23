package main

import "fmt"

// 2. Write a program in C to read n number of values in an array and display them in reverse order.
// Test Data :
// Input the number of elements to store in the array :3
// Input 3 number of elements in the array :
// element - 0 : 2
// element - 1 : 5
// element - 2 : 7
// Expected Output :
// The values stored into the array are :
// 2 5 7
// The values stored into the array in reverse are :
// 7 5 2

func main() {
	var n int
	fmt.Print("Enter the length of the array n=")
	fmt.Scanln(&n)
	var arr = make([]int, n)
	var arr_2 []int
	fmt.Printf("%T\n", arr)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter element %d arr[%d]=", i, i)
		fmt.Scanln(&arr[i])
	}
	fmt.Println("The values stored in the array are: \n", arr)
	for i := range arr {
		// arr_2=append(arr_2, arr[len(arr)-i-1])
		fmt.Println(arr[len(arr)-i-1])
	}
	fmt.Println("The values stored in the array in reverse are: \n", arr_2)
}

// Enter the length of the array n=56 67
// []int
// Enter element 0 arr[0]=1 Enter element 1 arr[1]=67
// Enter element 2 arr[2]=67
// Enter element 3 arr[3]=^Csignal: interrupt
