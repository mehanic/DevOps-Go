### **Explanation of the Code**
This Go program defines a function `sumOfUnique(nums []int) int` that calculates the sum of unique numbers in the given list (`nums`). A number is considered **unique** if it appears **only once** in the list.

---

### **Step-by-Step Explanation**
#### **1. Function Definition: `sumOfUnique(nums []int) int`**
```go
func sumOfUnique(nums []int) int {
    var summ = 0
    var x_map = make(map[int]int) // Create a map to store the frequency of each number
```
- We initialize `summ` (to store the final sum of unique numbers).
- We create a **map** (`x_map`) to store the frequency of each number in `nums`.

---

#### **2. Count Frequency of Each Number**
```go
for _, el := range nums {
    x_map[el]++
}
```
- We iterate through `nums`, incrementing the value in `x_map` for each occurrence of `el`.
- This keeps track of how many times each number appears in the list.

For example, if `nums = [1,2,3,2]`, after this loop, `x_map` will look like:

```go
map[1:1, 2:2, 3:1]
```
- `1` appears **once**.
- `2` appears **twice**.
- `3` appears **once**.

---

#### **3. Add Only Unique Elements**
```go
for key, val := range x_map {
    if val < 2 { // If a number appears less than 2 times (i.e., unique)
        summ += key
    }
}
```
- We iterate through `x_map`.
- If `val < 2` (meaning the number appears **only once**), we add its key to `summ`.

For our example:
- `1` → appears once → added to `summ`
- `2` → appears twice → **NOT** added
- `3` → appears once → added to `summ`

Thus, `summ = 1 + 3 = 4`.

---

#### **4. Return the Result**
```go
return summ
```
- After summing the unique numbers, we return `summ`.

---

### **5. Main Function**
```go
func main() {
    var nums = []int{1,2,3,2}
    fmt.Println(sumOfUnique(nums))
}
```
- `nums = [1,2,3,2]`
- Calls `sumOfUnique(nums)`
- Prints the output → `4`

---

### **Example Walkthrough**
#### **Input**
```go
nums = [1,2,3,2]
```
#### **Step 1: Count Elements**
```go
x_map = {1:1, 2:2, 3:1}
```
#### **Step 2: Sum Unique Elements**
- `1` is unique → added (`summ = 1`)
- `2` is **not** unique → skipped
- `3` is unique → added (`summ = 1 + 3 = 4`)

#### **Output**
```go
4
```

---

### **Time Complexity Analysis**
- **Building the frequency map** takes **O(n)**.
- **Iterating through the map** takes **O(n)**.
- **Total time complexity: O(n)**, which is efficient.

---

### **Final Thoughts**
- This program efficiently finds and sums unique elements in an array.
- It uses a **map (hash table)** to track occurrences.
- The approach ensures an **O(n) time complexity**, making it scalable for large inputs.

Would you like to modify this to handle more cases, like negative numbers or large inputs? 🚀