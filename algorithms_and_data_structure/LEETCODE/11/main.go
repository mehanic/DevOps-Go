package main
func firstUniqChar(s string) int {
    // Create a frequency map
    count := make(map[rune]int)

    // Count occurrences of each character
    for _, ch := range s {
        count[ch]++
    }

    // Find the first character with count == 1
    for i, ch := range s {
        if count[ch] == 1 {
            return i
        }
    }

    return -1
}
