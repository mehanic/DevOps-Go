Let's break down and explain the Go code in detail:

### Code Overview:

```go
package main

import "fmt"

func search(nums []int, target int) int {
    for i, el := range nums {
        if el == target {
            return i
        }
    }
    return -1
}

func main() {
    fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
```

### Explanation:

1. **`search` function**:
   ```go
   func search(nums []int, target int) int {
       for i, el := range nums {
           if el == target {
               return i
           }
       }
       return -1
   }
   ```
   - The `search` function takes two arguments:
     - `nums []int`: A slice of integers, which is the list to search through.
     - `target int`: The integer that you are trying to find within the slice.
   - **Inside the function**:
     - The `for` loop iterates over the `nums` slice using `range`. The `range` keyword gives two values for each iteration:
       - `i`: The index of the current element.
       - `el`: The value of the current element.
     - The `if` statement checks if the current element (`el`) is equal to the `target`. If it is, the index (`i`) of the current element is returned, which means the target has been found.
     - If the loop completes without finding the target, the function returns `-1`, indicating that the target is not present in the list.

2. **`main` function**:
   ```go
   func main() {
       fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
   }
   ```
   - The `main` function calls the `search` function with the arguments:
     - `nums = []int{-1, 0, 3, 5, 9, 12}` (the slice to search through)
     - `target = 9` (the target value to search for).
   - The result of the `search` function (either the index of the target or `-1`) is printed using `fmt.Println`.

### Example Walkthrough:

The `search` function is called with the list `[-1, 0, 3, 5, 9, 12]` and the target value `9`.

- **First Iteration**:
  - `i = 0`, `el = -1`
  - `el != target` (`-1 != 9`), so continue to the next iteration.
- **Second Iteration**:
  - `i = 1`, `el = 0`
  - `el != target` (`0 != 9`), so continue to the next iteration.
- **Third Iteration**:
  - `i = 2`, `el = 3`
  - `el != target` (`3 != 9`), so continue to the next iteration.
- **Fourth Iteration**:
  - `i = 3`, `el = 5`
  - `el != target` (`5 != 9`), so continue to the next iteration.
- **Fifth Iteration**:
  - `i = 4`, `el = 9`
  - `el == target` (`9 == 9`), so the function returns `4` (the index of `9` in the slice).

Thus, the function returns `4`, and the `fmt.Println` statement prints `4`.

### Final Output:
```
4
```

### Summary of the Logic:
- **Purpose**: This code is a simple implementation of linear search. It searches for a specific `target` in a slice of integers (`nums`).
- **Steps**:
  1. Iterate through the slice.
  2. If the target value is found, return the index.
  3. If the loop finishes and the target isn't found, return `-1` to indicate the target is not in the slice.
  
### Key Points:
- **Linear Search**: The approach used here is a **linear search**, which checks each element one by one until it finds the target or exhausts the list.
- **Return value**: If the target is found, it returns the index of the element; if not, it returns `-1`.
