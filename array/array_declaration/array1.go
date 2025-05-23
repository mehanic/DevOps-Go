package main

import "fmt"

func main() {
	var (
		myAge  byte
		herAge byte

		ages [2]byte

		tags [5]string

		zero [0]byte

		agesExp [2 + 1]byte
	)

	fmt.Printf("myAge             : %d\n", myAge)
	fmt.Printf("herAge            : %d\n", herAge)

	fmt.Printf("ages              : %#v\n", ages)
	fmt.Printf("tags              : %#v\n", tags)
	fmt.Printf("zero              : %#v\n", zero)
	fmt.Printf("agesExp           : %#v\n", agesExp)

	{
		var ages [2]int

		fmt.Println()
		fmt.Printf("ages              : %#v\n", ages)
		fmt.Printf("ages's type       : %T\n", ages)

		fmt.Println("len(ages)         :", len(ages))
		fmt.Println("ages[0]           :", ages[0])
		fmt.Println("ages[1]           :", ages[1])
		fmt.Println("ages[len(ages)-1] :", ages[len(ages)-1])

		ages[0] = 6
		ages[1] -= 3

		fmt.Println("ages[0]           :", ages[0])
		fmt.Println("ages[1]           :", ages[1])

		ages[0] *= ages[1]
		fmt.Println("ages[0]           :", ages[0])
		fmt.Println("ages[1]           :", ages[1])
	}
}

// myAge             : 0
// herAge            : 0
// ages              : [2]uint8{0x0, 0x0}
// tags              : [5]string{"", "", "", "", ""}
// zero              : [0]uint8{}
// agesExp           : [3]uint8{0x0, 0x0, 0x0}

// ages              : [2]int{0, 0}
// ages's type       : [2]int
// len(ages)         : 2
// ages[0]           : 0
// ages[1]           : 0
// ages[len(ages)-1] : 0
// ages[0]           : 6
// ages[1]           : -3
// ages[0]           : -18
// ages[1]           : -3
