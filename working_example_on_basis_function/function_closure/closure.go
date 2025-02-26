package closure

func closure(nums []int, sum func(l []int)) error {
	sum(nums)
	return nil
}
