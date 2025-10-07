package main

import "fmt"

func reverseString1(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseStringXOR(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i] ^= s[j]
		s[j] ^= s[i]
		s[i] ^= s[j]
	}
}

func reverseString(s []byte) {
	n := len(s)

	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}

func main() {
	k := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(k)
	fmt.Println(string(k))

	s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseStringXOR(s2)
	fmt.Println(string(s2))

	s3 := []byte{'B', 'e', 'r', 'l', 'i', 'n'}
	reverseString1(s3)
	fmt.Println(string(s3)) // Output: "hannaH"
}