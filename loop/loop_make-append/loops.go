package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("tutorial for loops")

	day := make([]string, 0)
	day = append(day, "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday")
	fmt.Println(day)

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Current generated day is", day[rand.Intn(7)])
	// for d := 0; d < len(day); d++ {
	// 	fmt.Println(day[d])
	// }

	// for i := range day {
	// fmt.Println(day[i])
	// }

	for index, value := range day {
		fmt.Println("Index is", index, " and value is", value)
	}

}


// tutorial for loops
// [Sunday Monday Tuesday Wednesday Thursday Friday Saturday]
// Current generated day is Monday
// Index is 0  and value is Sunday
// Index is 1  and value is Monday
// Index is 2  and value is Tuesday
// Index is 3  and value is Wednesday
// Index is 4  and value is Thursday
// Index is 5  and value is Friday
// Index is 6  and value is Saturday
