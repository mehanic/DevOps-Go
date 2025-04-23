package main

import "fmt"

// O(n) time | O(1) space
func IsValidSubsequence(array []int, sequence []int) bool {
	seqIdx := 0
	for _, value := range array {
		if seqIdx == len(sequence) {
			break
		}
		if value == sequence[seqIdx] {
			seqIdx += 1
		}
	}
	return seqIdx == len(sequence)

}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	fmt.Println(IsValidSubsequence(array, sequence))
	fmt.Println("--------------------")
	main1()
	main2()
	main3()
}

func main1() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, 10, -1}
	fmt.Println(IsValidSubsequence(array, sequence))

}

func main2() {
	array := []int{2, 4, 3, 7, 9, 5, 8}
	sequence := []int{4, 7, 5}
	fmt.Println(IsValidSubsequence(array, sequence))

}

func main3() {
	array := []int{2, 4, 3, 7, 9, 5, 8}
	sequence := []int{4, 7, 5}
	fmt.Println(IsValidSubsequence(array, sequence))
}
