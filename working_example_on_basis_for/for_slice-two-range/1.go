package main

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{4, 5, 6, 7}
	arr3 := [][]int{arr1, arr2}
	for _, v := range arr3 {
		for _, k := range v {
			println(k)
		}
	}
}
