package main

import "fmt"

func happyNumber(n int) bool {
	slow, fast := n, n
	for {
		slow = getNextNum(slow)
		fast = getNextNum(getNextNum(fast))
		if fast == 1 {
			return true
		} else if slow == fast {
			return false
		}
	}
}

func getNextNum(x int) int {
	nextNum := 0
	for x > 0 {
		digit := x % 10
		x /= 10
		nextNum += digit * digit
	}
	return nextNum
}

func main() {
	fmt.Println(happyNumber(19)) // true
	fmt.Println(happyNumber(2))  // false
}
