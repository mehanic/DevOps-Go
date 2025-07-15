func singleNumber(nums []int) int {
    for i := 0; i < len(nums); i++ {
        flag := true
        for j := 0; j < len(nums); j++ {
            if i != j && nums[i] == nums[j] {
                flag = false
                break
            }
        }
        if flag {
            return nums[i]
        }
    }
    return 0
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))         // Output: 1
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))   // Output: 4
	fmt.Println(singleNumber([]int{7}))               // Output: 7
}