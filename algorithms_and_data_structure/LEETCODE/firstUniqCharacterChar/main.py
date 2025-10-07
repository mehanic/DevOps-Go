from collections import Counter
def firstUniqChar(self, s: str) -> int:
        freq = Counter(s)  # Count occurrences of each character

        for i, ch in enumerate(s):
            if freq[ch] == 1:
                return i  # Return the index of the first non-repeating character

        return -1  # If no unique character is found


print(firstUniqChar("leetcode"))      # Output: 0
print(firstUniqChar("loveleetcode"))  # Output: 2
print(firstUniqChar("aabb"))          # Output: -1

