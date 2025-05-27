package main

func kSumSubarraysOptimized(nums []int, k int) int {
    count := 0
    prefixSumMap := map[int]int{0: 1}
    currPrefixSum := 0

    for _, num := range nums {
        currPrefixSum += num
        if freq, found := prefixSumMap[currPrefixSum - k]; found {
            count += freq
        }
        prefixSumMap[currPrefixSum]++
    }

    return count
}
