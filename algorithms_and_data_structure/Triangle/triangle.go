package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter 3 numbers a, b, and c to check if they can form a triangle")

	reader := bufio.NewReader(os.Stdin)

	// Enter a (must not be zero)
	fmt.Print("Enter a = ")
	str1, err := reader.ReadString('\n')
	str1 = strings.TrimSuffix(str1, "\n")
	a, err := strconv.ParseFloat(str1, 64)
	if a == 0 {
		fmt.Println(err.Error())
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Enter b
	fmt.Print("Enter b = ")
	str2, err := reader.ReadString('\n')
	str2 = strings.TrimSuffix(str2, "\n")
	b, err := strconv.ParseFloat(str2, 64)
	if b == 0 {
		fmt.Println(err.Error())
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Enter c
	fmt.Print("Enter c = ")
	str3, err := reader.ReadString('\n')
	str3 = strings.TrimSuffix(str3, "\n")
	c, err := strconv.ParseFloat(str3, 64)
	if c == 0 {
		fmt.Println(err.Error())
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if a+b > c && a+c > b && b+c > a {
		if a*a+b*b == c*c || b*b+c*c == a*a || c*c+a*a == b*b {
			fmt.Println("The three numbers you entered are the lengths of a right triangle")
		} else {
			if a == b && b == c {
				fmt.Println("The three numbers you entered are the lengths of an equilateral triangle")
			} else {
				if a*a > b*b+c*c || b*b > a*a+c*c || c*c > a*a+b*b {
					fmt.Println("The three numbers you entered are the lengths of an obtuse triangle")
				} else {
					fmt.Println("The three numbers you entered are the lengths of an acute triangle")
				}
			}
		}
	} else {
		fmt.Println("a = " + str1 + ", b = " + str2 + ", and c = " + str3 + " do not form the lengths of a triangle")
	}

	// You could also use switch-case instead of if-else...
}
