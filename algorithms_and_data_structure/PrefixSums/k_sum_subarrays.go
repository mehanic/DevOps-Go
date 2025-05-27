package main

func kSumSubarrays(nums []int, k int) int {
    n := len(nums)
    count := 0

    prefixSum := make([]int, n+1)
    prefixSum[0] = 0

    for i := 0; i < n; i++ {
        prefixSum[i+1] = prefixSum[i] + nums[i]
    }

    for j := 1; j <= n; j++ {
        for i := 1; i <= j; i++ {
            if prefixSum[j]-prefixSum[i-1] == k {
                count++
            }
        }
    }

    return count
}
