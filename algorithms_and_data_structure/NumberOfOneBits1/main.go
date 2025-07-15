func hammingWeight(n int) int {
	res := 0
	for i := 0; i < 32; i++ {
		if (1<<i)&n != 0 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(hammingWeight(11))     // Output: 3 (binary: 1011)
	fmt.Println(hammingWeight(128))    // Output: 1 (binary: 10000000)
	fmt.Println(hammingWeight(255))    // Output: 8 (binary: 11111111)
}