Let's break down the Go code and explain the logic step by step:

### Goal:
The function `lengthOfLastWord` is designed to return the length of the last word in a given string `s`. A word is defined as a sequence of non-space characters, and any leading or trailing spaces should be ignored. 

### Code Breakdown:

#### 1. `strings.TrimSpace(s)`
```go
s = strings.TrimSpace(s)
```
- **Purpose**: This line removes any leading and trailing whitespace characters from the input string `s`.
- **Example**: If the input is `"   fly me   to   the moon  "`, after this line, `s` will become `"fly me   to   the moon"`. This step ensures that we do not count leading or trailing spaces when processing the string.

#### 2. `var str = ""`
```go
var str = ""
```
- **Purpose**: We initialize an empty string `str`, which will be used to build the last word in reverse. This is important because we will iterate over the string from the end to the beginning, and appending characters to `str` will eventually hold the reversed last word.

#### 3. Iterate over the string in reverse:
```go
for i := len(s) - 1; i >= 0; i-- {
    if string(s[i]) != " " {
        str += string(s[i])
    } else {
        break
    }
}
```
- **Purpose**: The loop iterates through the string `s` from the last character to the first.
  - `i := len(s) - 1`: This sets the starting point of the loop to the last index of the string.
  - `i >= 0`: The loop continues until we reach the first character of the string.
  - `i--`: Decrements `i` to move through the string from right to left.

- Inside the loop:
  - `if string(s[i]) != " "`: This checks if the current character is not a space. If it's not a space, we append that character to the `str` variable. This builds the last word in reverse order.
  - `else { break }`: If we encounter a space, it means we've reached the end of the last word, so we exit the loop using `break`.

#### 4. Return the length of the last word:
```go
return len(str)
```
- **Purpose**: After the loop has collected all the characters of the last word (in reverse order), we return the length of `str`. The length of `str` will be the length of the last word, as `str` now contains the characters of the last word in reverse order.

#### 5. Main function:
```go
var s = "   fly me   to   the moon  "
fmt.Println(lengthOfLastWord(s))
```
- **Purpose**: The `main` function initializes a string `s` and calls the `lengthOfLastWord` function with this string as the argument. It then prints the result (the length of the last word).

### Example Walkthrough:

1. **Input**: `"   fly me   to   the moon  "`
2. **Step 1**: The string `s` is trimmed of leading and trailing spaces using `strings.TrimSpace(s)`. The string becomes `"fly me   to   the moon"`.
3. **Step 2**: The loop starts at the last character of the string (`s[len(s)-1]` which is `"n"`).
   - `"n"` is added to `str`.
   - Next, the loop moves to `"o"`, which is added to `str`.
   - The loop continues adding `"o"`, `"m"`, `" "`, but stops when it encounters a space `" "`.
4. **Step 3**: The string `str` now holds the last word in reverse: `"noom"`.
5. **Step 4**: The length of `"noom"` is `4`, so the function returns `4`.

### Output:
```go
4
```

### Key Points:
1. **Trim leading/trailing spaces**: `strings.TrimSpace(s)` is used to remove any unnecessary spaces from the beginning and end of the string, so they do not interfere with the word length calculation.
2. **Reverse traversal**: The loop iterates through the string from the end to the beginning, checking for the last word by adding non-space characters to `str`.
3. **Break on space**: When a space is encountered, the loop stops because we know that the last word has ended.

### Efficiency:
- **Time complexity**: O(n), where `n` is the length of the string. We are iterating over the string once to trim the spaces and then iterate again (at most once) to find the last word.
- **Space complexity**: O(1), as we only use a fixed amount of extra space (the variable `str` which stores the last word in reverse).

This approach is simple, efficient, and easy to understand!