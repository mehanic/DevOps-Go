
from typing import List


class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        n = len(nums)

        k = k % n

        def reverse(start, end):
            while start < end:
                nums[start], nums[end] = nums[end], nums[start]
                start += 1
                end -= 1

        # Reverse entire array
        reverse(0, n - 1)
        # Reverse first k elements
        reverse(0, k - 1)
        # Reverse remaining
        reverse(k, n - 1)


