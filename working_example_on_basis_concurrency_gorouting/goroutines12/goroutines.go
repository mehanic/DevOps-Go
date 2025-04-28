package main

import (
	"fmt"
	"time"
)

func Tek() {
	for i := 1; i < 10; i += 2 {

		fmt.Println("tek : ", i)
		time.Sleep(1 * time.Second)
	}
}
func Goroutines() {
	for i := 0; i < 10; i += 2 {
		fmt.Println("çift : ", i)
		time.Sleep(1 * time.Second)
	}

}

func main() {
	Tek()
	fmt.Println("----------")
	Goroutines()
}

// ─λ go run goroutines.go                                                                  1 (0.015s) < 15:18:09
// tek :  1
// tek :  3
// tek :  5
// tek :  7
// tek :  9
// ----------
// çift :  0
// çift :  2
// çift :  4
// çift :  6
// çift :  8
