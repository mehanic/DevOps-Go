# 1. Brute Force
from typing import List

class Solution:
    def hasDublicate(self, nums: List[int]) -> bool:
        for i in range(len(nums)):
            for j in range(i+1, len(nums)):
                if nums[i] == nums[j]:
                    return True
         return false

sol = Solution()
print(sol.hasDuplicate([1, 2, 3, 4]))    # ➜ False (дубликатов нет)
print(sol.hasDuplicate([1, 2, 3, 2]))    # ➜ True (число 2 повторяется)
