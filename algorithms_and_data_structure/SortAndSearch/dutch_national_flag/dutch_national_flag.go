package main

import "fmt"

func dutchNationalFlag(nums []int) {
	i, left, right := 0, 0, len(nums)-1
	for i <= right {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			i++
		}
	}
}


func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	dutchNationalFlag(nums)
	fmt.Println(nums) // Output: [0 0 1 1 2 2]
}