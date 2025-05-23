package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)
}

// sum the number verbose
func main1() {
	const min, max = 1, 10

	var sum int
	for i := min; i <= max; i++ {
		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}

// sum up to n
func main2() {
	if len(os.Args) != 3 {
		fmt.Println("gimme two numbers")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("wrong numbers")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}

// only events
func main3() {
	if len(os.Args) < 3 {
		fmt.Println("gimme two numbers")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("wrong numbers")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i

		fmt.Print(i)
		if i < max-1 {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}

// break-up
func main4() {
	if len(os.Args) < 3 {
		fmt.Println("gimme two numbers")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("wrong numbers")
		return
	}

	var (
		i   = min
		sum int
	)

	for {
		if i > max {
			break
		} else if i%2 != 0 {
			i++
			continue
		}

		fmt.Print(i)
		if i < max-1 {
			fmt.Print(" + ")
		}

		sum += i
		i++
	}
	fmt.Printf(" = %d\n", sum)
}

// infinite-kill
func main5() {
	for {
		var c string

		switch rand.Intn(4) {
		case 0:
			c = "\\"
		case 1:
			c = "/"
		case 2:
			c = "|"
		case 3:
			c = "-"
		}
		fmt.Printf("\r%s Please Wait. Processing....", c)
		time.Sleep(time.Millisecond * 250)
	}
}

func main6() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Give me the size of the table")
		return
	}

	size, err := strconv.Atoi(args[1])
	if err != nil || size <= 0 {
		fmt.Println("Wrong size")
		return
	}

	// print the header
	fmt.Printf("%5s", "X")
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= size; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= size; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}

//math-table

const (
	validOps = "%*/+-"

	usageMsg       = "Usage: [op=" + validOps + "] [size]"
	sizeMissingMsg = "Size is missing"
	invalidOpMsg   = `Invalid operator.
Valid ops one of: ` + validOps

	invalidOp = -1
)

func main7() {
	args := os.Args[1:]

	switch l := len(args); {
	case l == 1:
		fmt.Println(sizeMissingMsg)
		fallthrough
	case l < 1:
		fmt.Println(usageMsg)
		return
	}

	op := args[0]
	if strings.IndexAny(op, validOps) == invalidOp {
		fmt.Println(invalidOpMsg)
		return
	}

	size, err := strconv.Atoi(args[1])
	if err != nil || size <= 0 {
		fmt.Println("Wrong size")
		return
	}

	fmt.Printf("%5s", op)
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= size; j++ {
			var res int

			switch op {
			case "*":
				res = i * j
			case "%":
				if j != 0 {
					res = i % j
				}
			case "/":
				if j != 0 {
					res = i / j
				}
			case "+":
				res = i + j
			case "-":
				res = i - j
			}

			fmt.Printf("%5d", res)
		}
		fmt.Println()
	}
}

//
