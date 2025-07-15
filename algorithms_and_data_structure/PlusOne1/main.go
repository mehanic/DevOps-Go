func plusOne(digits []int) []int {
    if len(digits) == 0 {
        return []int{1}
    }

    if digits[len(digits)-1] < 9 {
        digits[len(digits)-1]++
        return digits
    } else {
        return append(plusOne(digits[:len(digits)-1]), 0)
    }
}


func main() {
	fmt.Println(plusOne([]int{1, 2, 3})) // Output: [1 2 4]
	fmt.Println(plusOne([]int{4, 3, 9})) // Output: [4 4 0]
	fmt.Println(plusOne([]int{9, 9, 9})) // Output: [1 0 0 0]
}