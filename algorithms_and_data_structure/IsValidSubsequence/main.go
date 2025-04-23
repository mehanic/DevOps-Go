package main

import "fmt"

func IsValidSubsequence(array []int, sequence []int) bool {
	arrIdx := 0 // Указатель на элементы в array
	seqIdx := 0 // Указатель на элементы в sequence

	// Перебираем массив, пока не дойдем до конца одного из них
	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] { // Если элементы совпадают
			seqIdx += 1 // Перемещаем указатель sequence на следующий элемент
		}
		arrIdx += 1 // Всегда перемещаем указатель array
	}

	// Если seqIdx дошел до конца sequence, значит, все элементы были найдены
	return seqIdx == len(sequence)
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	fmt.Println(IsValidSubsequence(array, sequence)) // true
	fmt.Println("--------------------")
	main1()

}

func main1() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, 10, -1}
	fmt.Println(IsValidSubsequence(array, sequence)) // false

}
