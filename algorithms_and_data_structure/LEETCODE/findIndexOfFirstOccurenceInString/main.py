def strStr(haystack: str, needle: str) -> int:
    if not needle:
        return 0
    
    for i in range(len(haystack) - len(needle) + 1):
        # Check substring from i to i+len(needle)
        if haystack[i:i+len(needle)] == needle:
            return i
    return -1

# Examples
print(strStr("sadbutsad", "sad"))  # Output: 0
print(strStr("leetcode", "leeto")) # Output: -1
