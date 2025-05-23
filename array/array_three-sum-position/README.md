The Go program implements a solution to the **3-sum problem**, where the goal is to find all unique triplets in an array that sum to zero. Let's break down the program, step by step, to understand the logic and how it works.

### **The `threeSum` Function:**

```go
func threeSum(nums []int) [][]int {
    res := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        l, r := i+1, len(nums)-1
        for l < r {
            s := nums[i] + nums[l] + nums[r]
            if s > 0 {
                r -= 1
            } else if s < 0 {
                l += 1
            } else {
                res = append(res, []int{nums[i], nums[l], nums[r]})
                for l < r && nums[l] == nums[l+1] {
                    l += 1
                }
                for l < r && nums[r] == nums[r-1] {
                    r -= 1
                }
                l += 1
                r -= 1
            }
        }
    }
    return res
}
```

### **Explanation of the Code:**

1. **Sorting the Input Array:**
   ```go
   sort.Ints(nums)
   ```
   - The first step is sorting the input array. Sorting helps in efficiently finding triplets because it allows the use of the **two-pointer technique**. Once the array is sorted, if you are looking for a sum of zero, the problem becomes much simpler because you can systematically check the sum of pairs and adjust the pointers accordingly.

2. **Outer Loop:**
   ```go
   for i := 0; i < len(nums)-2; i++ {
       if i > 0 && nums[i] == nums[i-1] {
           continue
       }
   ```
   - The outer loop iterates through each element in the array, except for the last two (because we need at least three elements to form a triplet). 
   - `i > 0 && nums[i] == nums[i-1]` is used to **skip duplicates**. If the current element `nums[i]` is the same as the previous one, the loop skips this iteration to avoid checking the same triplet multiple times.

3. **Two-Pointer Technique:**
   ```go
   l, r := i+1, len(nums)-1
   for l < r {
       s := nums[i] + nums[l] + nums[r]
       if s > 0 {
           r -= 1
       } else if s < 0 {
           l += 1
       } else {
           res = append(res, []int{nums[i], nums[l], nums[r]})
           for l < r && nums[l] == nums[l+1] {
               l += 1
           }
           for l < r && nums[r] == nums[r-1] {
               r -= 1
           }
           l += 1
           r -= 1
       }
   }
   ```
   - After sorting, we use two pointers (`l` and `r`) to find pairs of numbers that, when added to `nums[i]`, sum to zero.
   - **Left pointer (`l`)** starts from the element right after `i`, and **right pointer (`r`)** starts from the end of the array.
   - The sum `s` of `nums[i]`, `nums[l]`, and `nums[r]` is computed:
     - If `s > 0`, the sum is too large, so we move the **right pointer (`r`) left** to decrease the sum.
     - If `s < 0`, the sum is too small, so we move the **left pointer (`l`) right** to increase the sum.
     - If `s == 0`, we have found a valid triplet, and we add it to the result list `res`.

4. **Skipping Duplicates:**
   ```go
   for l < r && nums[l] == nums[l+1] {
       l += 1
   }
   for l < r && nums[r] == nums[r-1] {
       r -= 1
   }
   ```
   - After finding a valid triplet, we need to **skip duplicates** to avoid adding the same triplet again.
   - The inner loops increment `l` (left pointer) and decrement `r` (right pointer) while the adjacent elements are the same as the current ones.

5. **Updating Pointers:**
   ```go
   l += 1
   r -= 1
   ```
   - After handling duplicates, we move the pointers one step forward and backward to continue searching for new potential triplets.

### **Return Result:**
```go
return res
```
- Finally, after the outer loop finishes, the function returns the result `res`, which contains all the unique triplets that sum to zero.

---

### **Main Function:**

```go
func main() {
    nums := []int{-1, 0, 1, 2, -1, -4}
    fmt.Println("Input array:", nums)

    result := threeSum(nums)

    fmt.Println("Unique triplets that sum to zero:", result)
}
```

1. **Input Array:**
   - The input array is `nums := []int{-1, 0, 1, 2, -1, -4}`.
   
2. **Function Call:**
   - The `threeSum` function is called with the input array.

3. **Output:**
   - The program prints the input array and the list of unique triplets that sum to zero. The result is:
   ```go
   Unique triplets that sum to zero: [[-1 -1 2] [-1 0 1]]
   ```

### **Example Walkthrough:**

For the input array `[-1, 0, 1, 2, -1, -4]`:

1. **Step 1: Sorting the Array:**
   - The array is sorted: `[-4, -1, -1, 0, 1, 2]`.

2. **Step 2: Outer Loop Iteration (i = 0):**
   - The first number is `-4`. The two pointers `l` and `r` start at `1` and `5` respectively.
   - The sum is `-4 + (-1) + 2 = -3` (too small), so the left pointer `l` is incremented.
   - The sum is `-4 + 0 + 2 = -2` (still too small), so the left pointer `l` is incremented.
   - The sum is `-4 + 1 + 2 = -1` (still too small), so the left pointer `l` is incremented.
   - The sum is `-4 + 1 + 2 = -1`, and the left pointer `l` reaches `2` and `r` reaches `4`, this will be time