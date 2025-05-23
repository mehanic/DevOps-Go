package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is a lesson on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}

	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Pineapple", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 867
	//highScores[4] = 777 ___NOTE: Trying to allocate new data into the highscore slice-array as declared will throw an error as we initialized the slice-array with default memory allocation of 4. Thus, to add data to the slice, we use the append method. This will allow the allocation of new memory.

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	//How to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
