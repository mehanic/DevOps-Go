func countBits(n int) []int {
    res := make([]int, n+1)
    for num := 0; num <= n; num++ {
        one := 0
        for i := 0; i < 32; i++ {
            if num&(1<<i) != 0 {
                one++
            }
        }
        res[num] = one
    }
    return res
}


func main() {
	result := countBits(5)
	fmt.Println(result) // Output: [0 1 1 2 1 2]
}