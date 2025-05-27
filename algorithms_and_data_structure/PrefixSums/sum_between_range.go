package main

type SumBetweenRange struct {
	prefixSum []int
}

func NewSumBetweenRange(nums []int) *SumBetweenRange {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}
	return &SumBetweenRange{prefixSum: prefixSum}
}

func (s *SumBetweenRange) SumRange(i, j int) int {
	if i == 0 {
		return s.prefixSum[j]
	}
	return s.prefixSum[j] - s.prefixSum[i-1]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sbr := NewSumBetweenRange(nums)
	println(sbr.SumRange(1, 3)) // 2+3+4=9
}
