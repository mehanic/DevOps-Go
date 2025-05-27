package main

import (
	"fmt"
)

func swapOddAndEvenBits(n int) int {
	evenMask := 0x55555555 // 01010101010101010101010101010101
	oddMask := 0xAAAAAAAA  // 10101010101010101010101010101010

	evenBits := n & evenMask
	oddBits := n & oddMask

	// Сдвигаем четные биты влево, нечетные — вправо, объединяем результат
	return (evenBits << 1) | (oddBits >> 1)
}

func main() {
	n := 23 // Пример: 00010111
	fmt.Printf("Result: %d\n", swapOddAndEvenBits(n))
}
