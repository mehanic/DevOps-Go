type KthLargest struct {
    k   int
    arr []int
}

func Constructor(k int, nums []int) KthLargest {
    return KthLargest{k: k, arr: nums}
}

func (this *KthLargest) Add(val int) int {
    this.arr = append(this.arr, val)
    sort.Ints(this.arr)
    return this.arr[len(this.arr)-this.k]
}

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	kthLargest := Constructor(k, nums)

	fmt.Println(kthLargest.Add(3))  // -> 4
	fmt.Println(kthLargest.Add(5))  // -> 5
	fmt.Println(kthLargest.Add(10)) // -> 5
	fmt.Println(kthLargest.Add(9))  // -> 8
	fmt.Println(kthLargest.Add(4))  // -> 8
}