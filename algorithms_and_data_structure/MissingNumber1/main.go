func missingNumber(nums []int) int {
    n := len(nums)
    sort.Ints(nums)
    for i := 0; i < n; i++ {
        if nums[i] != i {
            return i
        }
    }
    return n
}

func main() {
    example1 := []int{3, 0, 1}
    example2 := []int{0, 1}
    example3 := []int{9,6,4,2,3,5,7,0,1}

    fmt.Println(missingNumber(example1)) // Output: 2
    fmt.Println(missingNumber(example2)) // Output: 2
    fmt.Println(missingNumber(example3)) // Output: 8
}