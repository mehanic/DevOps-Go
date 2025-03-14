Let's go through the code step by step.

### Code Overview:
```go
package main

import "fmt"

func removeElement(nums []int, val int) int {
    fmt.Println(len(nums))
    var count = 0
    for _, el := range nums {
        if el != val {
            // fmt.Println(el, nums)
            nums[count] = el
            count++
        }
    }
    return count
}

func main() {
    fmt.Println(removeElement([]int{3, 2, 2, 4, 5, 6, 3, 6, 7, 8, 3}, 2))
}
```

### Explanation:

#### `removeElement` Function:
```go
func removeElement(nums []int, val int) int {
    fmt.Println(len(nums))
    var count = 0
    for _, el := range nums {
        if el != val {
            // fmt.Println(el, nums)
            nums[count] = el
            count++
        }
    }
    return count
}
```

- **Parameters**:
  - `nums []int`: A slice of integers representing the list of numbers.
  - `val int`: The value that needs to be removed from the `nums` slice.

- **Purpose**:
  - The function is designed to remove all occurrences of the number `val` from the `nums` slice.
  - The goal is to return the new length of the slice after removing `val`, with the remaining elements (those that are not equal to `val`) shifted to the left.
  
- **Steps**:
  1. **Print the length of `nums`**:
     ```go
     fmt.Println(len(nums))
     ```
     This prints the length of the original `nums` slice before any modifications are made.
  
  2. **Variable Initialization**:
     ```go
     var count = 0
     ```
     A variable `count` is initialized to `0`. This will be used to keep track of the index where the next element (that is not equal to `val`) should be placed.

  3. **Loop through `nums`**:
     ```go
     for _, el := range nums {
         if el != val {
             nums[count] = el
             count++
         }
     }
     ```
     - The `for _, el := range nums` loop iterates over each element (`el`) in the `nums` slice.
     - If the element `el` is **not** equal to `val`, the element is placed at the `count` index in the `nums` slice.
     - The `count` is then incremented to point to the next available position.
     
     Essentially, this loop removes all occurrences of `val` by shifting the remaining elements to the left.

  4. **Return the new length**:
     ```go
     return count
     ```
     After the loop finishes, `count` will hold the number of elements that are not equal to `val`, which is the new length of the `nums` slice. This is returned as the result of the function.

#### `main` Function:
```go
func main() {
    fmt.Println(removeElement([]int{3, 2, 2, 4, 5, 6, 3, 6, 7, 8, 3}, 2))
}
```

- This function calls `removeElement` with the `nums` slice `{3, 2, 2, 4, 5, 6, 3, 6, 7, 8, 3}` and the value `2` to be removed.
- The result of `removeElement` (the new length of the slice after removing `2`s) is printed.

### Detailed Example:

Let's walk through the function call `removeElement([]int{3, 2, 2, 4, 5, 6, 3, 6, 7, 8, 3}, 2)`.

1. **Initial state**: The input slice is `{3, 2, 2, 4, 5, 6, 3, 6, 7, 8, 3}` and the value to remove is `2`.
   
2. **Iteration 1**: The first element is `3`. Since `3` is not equal to `2`, we move it to index `0` (which is already where it is). `count` becomes `1`.

3. **Iteration 2**: The second element is `2`. Since `2` is equal to `val`, we skip it and do not move it.

4. **Iteration 3**: The third element is `2`. It is also skipped because it's equal to `val`.

5. **Iteration 4**: The fourth element is `4`. Since `4` is not equal to `2`, we move it to index `1` (replacing `2`), and `count` becomes `2`.

6. **Iteration 5**: The fifth element is `5`. We move `5` to index `2` (replacing the previous `2`), and `count` becomes `3`.

7. **Iteration 6**: The sixth element is `6`. We move `6` to index `3`, and `count` becomes `4`.

8. **Iteration 7**: The seventh element is `3`. We move `3` to index `4`, and `count` becomes `5`.

9. **Iteration 8**: The eighth element is `6`. We move `6` to index `5`, and `count` becomes `6`.

10. **Iteration 9**: The ninth element is `7`. We move `7` to index `6`, and `count` becomes `7`.

11. **Iteration 10**: The tenth element is `8`. We move `8` to index `7`, and `count` becomes `8`.

12. **Iteration 11**: The eleventh element is `3`. We move `3` to index `8`, and `count` becomes `9`.

After all iterations, the modified `nums` slice is `{3, 4, 5, 6, 3, 6, 7, 8, 3, _, _, _}` (the trailing elements are not important, as they are beyond the new length).

### Final Output:
1. The function returns the new length of the `nums` slice after removing `2`, which is `9`.
2. The `main` function prints the result:
   ```
   9
   ```

### Summary of Key Points:
- **Goal**: The `removeElement` function removes all occurrences of a specified value (`val`) from the `nums` slice and returns the new length of the slice.
- **Approach**: It uses a single loop to shift the elements that are not equal to `val` to the front of the slice.
- **Time Complexity**: The time complexity of this approach is **O(n)**, where `n` is the length of the `nums` slice, as it iterates over the slice once.
