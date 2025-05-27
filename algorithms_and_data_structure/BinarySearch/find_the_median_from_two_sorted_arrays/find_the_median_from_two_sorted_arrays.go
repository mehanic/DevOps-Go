package main

import (
	"fmt"
	"math"
)

func findTheMedianFromTwoSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array
	if len(nums2) < len(nums1) {
		return findTheMedianFromTwoSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	halfTotalLen := (m + n) / 2
	left, right := 0, m-1

	for {
		L1Index := (left + right) / 2
		L2Index := halfTotalLen - (L1Index + 1) - 1

		L1 := math.Inf(-1)
		if L1Index >= 0 {
			L1 = float64(nums1[L1Index])
		}

		R1 := math.Inf(1)
		if L1Index+1 < m {
			R1 = float64(nums1[L1Index+1])
		}

		L2 := math.Inf(-1)
		if L2Index >= 0 {
			L2 = float64(nums2[L2Index])
		}

		R2 := math.Inf(1)
		if L2Index+1 < n {
			R2 = float64(nums2[L2Index+1])
		}

		if L1 > R2 {
			right = L1Index - 1
		} else if L2 > R1 {
			left = L1Index + 1
		} else {
			if (m+n)%2 == 0 {
				return (math.Max(L1, L2) + math.Min(R1, R2)) / 2.0
			} else {
				return math.Min(R1, R2)
			}
		}
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findTheMedianFromTwoSortedArrays(nums1, nums2)) // Output: 2.0
}
