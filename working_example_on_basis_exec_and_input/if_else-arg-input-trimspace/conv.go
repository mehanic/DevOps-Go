package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a rating between 0 to 5")
	input, _ := reader.ReadString('\n')
	fmt.Println("Your Input is ", input)
	fmt.Println("changing scale to 0 to 100")
	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
	newRating = newRating * 20
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New rating is ", newRating)
	}
}


// Enter a rating between 0 to 5
// 3
// Your Input is  3

// changing scale to 0 to 100
// New rating is  60
