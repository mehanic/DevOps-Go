The provided Go program implements a function called `summaryRanges` that takes a sorted array of integers and returns a list of strings, where each string represents a range of consecutive numbers in the array. Let's break down the program to explain the rules and logic behind it.

### **1. The `summaryRanges` Function:**

```go
func summaryRanges(nums []int) []string {
    res := []string{}
    if len(nums) == 1 {
        res = append(res, strconv.Itoa(nums[0]))
        return res
    }
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        for i+1 < len(nums) && nums[i+1]-nums[i] == 1 {
            i += 1
        }
        if nums[i] != num {
            var tempBuffer bytes.Buffer
            tempBuffer.WriteString(strconv.Itoa(num))
            tempBuffer.WriteString("->")
            tempBuffer.WriteString(strconv.Itoa(nums[i]))
            res = append(res, tempBuffer.String())
        } else {
            res = append(res, strconv.Itoa(num))
        }
    }
    return res
}
```

### **Explanation:**

- **Input:** The function `summaryRanges` takes a sorted slice of integers `nums`.

- **Output:** The function returns a slice of strings, where each string represents a range of consecutive numbers, formatted as `"start->end"`. If a number is not part of a range, it is returned as just that number.

#### **Step-by-step Explanation of the Code:**

1. **Edge Case for Single Element:**
   ```go
   if len(nums) == 1 {
       res = append(res, strconv.Itoa(nums[0]))
       return res
   }
   ```
   - This checks if the `nums` array has only one element. If so, it simply converts that element into a string and appends it to the result (`res`). It then returns the result.
   - This is an edge case to handle arrays with only one number, where no range is possible.

2. **Main Loop:**
   ```go
   for i := 0; i < len(nums); i++ {
       num := nums[i]
       for i+1 < len(nums) && nums[i+1]-nums[i] == 1 {
           i += 1
       }
   ```
   - This outer loop iterates through the `nums` slice, and for each element, it stores the current element in `num`.
   - The inner `for` loop checks if the next element (`nums[i+1]`) is exactly 1 greater than the current element (`nums[i]`). If true, it increments the index `i` to continue checking consecutive elements in the sequence.

3. **Detecting a Range or Single Number:**
   ```go
   if nums[i] != num {
       var tempBuffer bytes.Buffer
       tempBuffer.WriteString(strconv.Itoa(num))
       tempBuffer.WriteString("->")
       tempBuffer.WriteString(strconv.Itoa(nums[i]))
       res = append(res, tempBuffer.String())
   } else {
       res = append(res, strconv.Itoa(num))
   }
   ```
   - After exiting the inner loop, the code checks if the current element (`nums[i]`) is different from the starting element (`num`), which means we have found a range of consecutive numbers.
   - If a range is detected, it formats the range as `"start->end"`, where `num` is the starting number and `nums[i]` is the last number in the range.
   - If there was no range (i.e., the number is a standalone value), it simply appends the number as a string to the result slice.

4. **Return the Result:**
   ```go
   return res
   ```
   - Finally, the function returns the result slice `res`, which contains the formatted string representations of the ranges.

### **2. Example Walkthrough:**

For the input `nums := []int{0, 1, 2, 4, 5, 7}`, the function works as follows:

- **Iteration 1 (`i = 0`):**
  - `num = 0`, checks if `nums[1] - nums[0] == 1` (i.e., `1 - 0 == 1`). This is true, so increment `i` to `1`.
  - Checks if `nums[2] - nums[1] == 1` (i.e., `2 - 1 == 1`). This is true, so increment `i` to `2`.
  - Now, `nums[2] != num`, so we have a range from `0` to `2`, and the string `"0->2"` is appended to `res`.

- **Iteration 2 (`i = 3`):**
  - `num = 4`, checks if `nums[4] - nums[3] == 1` (i.e., `5 - 4 == 1`). This is true, so increment `i` to `4`.
  - Now, `nums[4] != num`, so we have a range from `4` to `5`, and the string `"4->5"` is appended to `res`.

- **Iteration 3 (`i = 5`):**
  - `num = 7`, but there are no consecutive numbers after it, so `"7"` is appended as a standalone number.

Final result:
```
res = ["0->2", "4->5", "7"]
```

### **3. Output:**

```go
Input array: [0 1 2 4 5 7]
Summary of ranges: [0->2 4->5 7]
```

### **4. Summary of Key Rules:**

1. **Range Detection:**
   - If a number is followed by a consecutive number (difference of 1), it forms a range.
   - If there’s a break in the sequence, the previous range is finalized and a new range or number is considered.

2. **Appending Results:**
   - If a range is found, it is formatted as `"start->end"`.
   - If a number is not part of a range, it is simply added as `"num"`.

3. **Buffer for String Construction:**
   - `bytes.Buffer` is used for efficiently concatenating strings (though this is not strictly necessary for this small-scale example).

### **5. Complexity:**

- **Time Complexity:** O(n), where n is the length of the input slice `nums`. Each element is checked once and only once.
- **Space Complexity:** O(n) for storing the result strings.

This solution is efficient and handles the problem of summarizing consecutive ranges in a sorted integer array effectively.