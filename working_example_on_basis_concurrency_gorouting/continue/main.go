package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	scores := []int{40, 50, 13, 90, 80, 60, 12, 39, 21, 60}
	fmt.Println("Old scores: ", scores)
MAKEUP:
	for index, score := range scores {
		if score > 40 {
			continue
		}

		rand.Seed(time.Now().UnixNano())
		scores[index] += rand.Intn(10)

		goto MAKEUP
	}

	fmt.Println("New scores: ", scores)
}

// ─λ go run main.go                                                                        0 (0.001s) < 13:56:59
// Old scores:  [40 50 13 90 80 60 12 39 21 60]
// New scores:  [47 50 48 90 80 60 45 48 46 60]
