package main

import (
	"fmt"
	"math"
)

func main() {
	var names [3]string
	fmt.Println(names)
	names[0] = "Jordan"
	names[2] = "James"
	fmt.Println(names)
	names[1] = "Em"
	fmt.Println(names)

	names2 := [3]string{"John, Ka, Li"}
	fmt.Println(names2, len(names2))

	fmt.Println("------------")
	for i := 0; i < len(names2); i++ {
		fmt.Println(names2[i])
	}

	marks := [5]float64{50, 50, 45, 39, 50}
	var sum float64 = 0

	for i := 0; i < len(marks); i++ {
		sum += marks[i]
	}
	fmt.Println(math.Round(sum / float64(len(marks))))
}

// [  ]
// [Jordan  James]
// [Jordan Em James]
// [John, Ka, Li  ] 3
// John, Ka, Li

// 47
