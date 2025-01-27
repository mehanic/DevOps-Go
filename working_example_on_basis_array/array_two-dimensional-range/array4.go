package main

import (
	"fmt"
)

func main() {
	// ------------------------------------
	// #1 - THE BEST WAY
	// ------------------------------------

	students := [...][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}

	var sum float64

	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
		}
	}

	const N = float64(len(students) * len(students[0]))
	fmt.Printf("Avg Grade: %g\n", sum/N)

	// ------------------------------------
	// #2 - SO SO WAY
	// ------------------------------------
	// fmt.Println("---------------------------")
	// // // You don't need to define the types for the inner arrays
	// students1 := [2][3]float64{
	// 	[3]float64{5, 6, 1},
	// 	[3]float64{9, 8, 4},
	// }

	// var sum1 float64

	// sum += students1[0][0] + students1[0][1] + students1[0][2]
	// sum += students1[1][0] + students1[1][1] + students1[1][2]

	// const D = float64(len(students1) * len(students1[0]))
	// fmt.Printf("Avg Grade: %g\n", sum1/D)

	// ------------------------------------
	// #3 - MANUAL WAY
	// ------------------------------------

	// student1 := [3]float64{5, 6, 1}
	// student2 := [3]float64{9, 8, 4}

	// var sum float64
	// sum += student1[0] + student1[1] + student1[2]
	// sum += student2[0] + student2[1] + student2[2]

	// const N = float64(len(student1) * 2)
	// fmt.Printf("Avg Grade: %g\n", sum/N)
}

// Avg Grade: 5.5
