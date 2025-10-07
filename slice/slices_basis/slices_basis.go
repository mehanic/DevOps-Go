package main

import "fmt"

func main() {
	animalArr := [4]string{
		"dog",
		"cat",
		"lie",
		"fox",
	}

	var a []string = animalArr[0:2]
	fmt.Println(a)
	b := animalArr[2:4]
	fmt.Println(b)
	b[0] = "dd"
	fmt.Println(b)
	// slice обердтка над массивом
	// animalSlices := []string{
	// 	"dog",
	// 	"cat",
	// 	"lie",
	// 	"fox",
	// }

	// чем массивы отличаются от слайсов в golang

	f := []string{"as", "d"}

	fmt.Println(f)
	arrays()
	slices()
}

func arrays() {
	// arrays
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	numbers := [...]int{1, 2, 3}
	fmt.Println(len(numbers))
}

func slices() {
	// slices
	letters := []string{"a", "b"}
	letters[0] = "f"
	letters = append(letters, "c", "d")
	fmt.Println(letters)
	createSlice := make([]string, 3)
	fmt.Println(len(createSlice))
	fmt.Println(cap(createSlice))
	createSlice[0] = "x"
	createSlice[1] = "z"
	createSlice[2] = "y"

	createSlice = append(createSlice, "g")
	fmt.Println(createSlice)
	// слайсы это массивы динамической длины!!!или капасту
	// капаситу увелчится вдвое когда длина слайса превышает макс.возможную капаситу
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
}

// [dog cat]
// [lie fox]
// [dd fox]
// [as d]
