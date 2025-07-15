package main

import (
	"fmt"
	"math"
)

// towerOfHanoi prints the initial setup and starts the recursive solution
func towerOfHanoi(numberOfDisks int) {
	fmt.Println("Tower of Hanoi \n")
	fmt.Println("S = Source tower\nI = Intermediate tower\nD = Destination tower\n")
	fmt.Printf("Total Disks in Source tower = %d\n\n", numberOfDisks)
	fmt.Printf("Solution (with minimum number of moves %.0f):\n\n", math.Pow(2, float64(numberOfDisks))-1)

	recursiveTowerOfHanoi(numberOfDisks, "S", "D", "I")
}

// recursiveTowerOfHanoi performs the actual recursive moves
func recursiveTowerOfHanoi(n int, source, destination, intermediate string) {
	if n == 1 {
		fmt.Printf("Disk %d moved from %s to %s\n", n, source, destination)
		return
	}
	recursiveTowerOfHanoi(n-1, source, intermediate, destination)
	fmt.Printf("Disk %d moved from %s to %s\n", n, source, destination)
	recursiveTowerOfHanoi(n-1, intermediate, destination, source)
}

func main() {
	numberOfDisks := 5
	towerOfHanoi(numberOfDisks)
}
