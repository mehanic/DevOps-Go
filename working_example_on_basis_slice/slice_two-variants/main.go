package main

import (
	"fmt"
)

func main() {
	gender := [2]string{"Male", "Female"}
	students := [8]string{"John", "James", "Ken", "Sam", "Peace", "Lucky", "Airon", "Charles"}

	fmt.Println("defined genders ", gender)

	fmt.Println("defined genders at index 1 ", gender[1])

	fmt.Println("sub slice  ", students[0:5])

	arrDescription := [4]string{"Learning", "is", "best", "things"}
	println(arrDescription)

}
