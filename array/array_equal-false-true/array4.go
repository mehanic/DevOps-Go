package main

import (
	"fmt"
)

func main() {
	var (
		// equal (types + elements are identical)::
		blue = [3]int{6, 9, 3}
		red  = [3]int{6, 9, 3}

		// equal (types + elements are identical):
		blue1 = [...]int{6, 9, 3}
		red1  = [3]int{6, 9, 3}

		// not equal (element ordering are different):
		blue2 = [3]int{6, 9, 3}
		red2  = [3]int{3, 9, 6}

		// not equal (the last elements are not equal):
		blue3 = [3]int{6, 9}
		red3  = [3]int{6, 9, 3}

		// not comparable (type mismatch: length):
		// blue4 = [3]int{6, 9, 3}
		// red4  = [5]int{6, 9, 3}

		// not comparable (type mismatch: element type):
		// blue5 = [3]int64{6, 9, 3}
		// red5  = [3]int{6, 9, 3}
	)
	fmt.Printf("blue bookcase : %v\n", blue)
	fmt.Printf("red bookcase  : %v\n", red)
	fmt.Println("Are they equal?", blue == red)

	fmt.Printf("blue bookcase : %v\n", blue1)
	fmt.Printf("red bookcase  : %v\n", red1)
	fmt.Println("Are they equal?", blue1 == red1)

	fmt.Printf("blue bookcase : %v\n", blue2)
	fmt.Printf("red bookcase  : %v\n", red2)
	fmt.Println("Are they equal?", blue2 == red2)

	fmt.Printf("blue bookcase : %v\n", blue3)
	fmt.Printf("red bookcase  : %v\n", red3)
	fmt.Println("Are they equal?", blue3 == red3)

	// fmt.Printf("blue bookcase : %v\n", blue5)
	// fmt.Printf("red bookcase  : %v\n", red5)
	// fmt.Println("Are they equal?", blue5 == red5)
}

// blue bookcase : [6 9 3]
// red bookcase  : [6 9 3]
// Are they equal? true
// blue bookcase : [6 9 3]
// red bookcase  : [6 9 3]
// Are they equal? true
// blue bookcase : [6 9 3]
// red bookcase  : [3 9 6]
// Are they equal? false
// blue bookcase : [6 9 0]
// red bookcase  : [6 9 3]
// Are they equal? false
