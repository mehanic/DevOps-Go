class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        return sorted(s) == sorted(t)

s = "listen"
t = "silent"

sol = Solution()
print(sol.isAnagram(s, t))  # Output: True
