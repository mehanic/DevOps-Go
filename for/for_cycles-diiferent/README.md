This Go code demonstrates different types of **`for` loops**, **conditional statements**, and **iteration over slices and matrices**. Let's go through the code step by step:

---

## **1. Loop through numbers from 1 to 100 and print numbers divisible by 3:**
```go
for i := 1; i <= 100; i++ {
	if i%3 == 0 {
		fmt.Println(i)
	}
}
```
- A basic `for` loop that iterates from `1` to `100`.
- The condition `if i % 3 == 0` checks if the number is divisible by 3.
- If true, the number is printed.

### ✅ **Output (partial):**
```
3
6
9
12
...
99
```

---

## **2. Skip numbers that are NOT divisible by 7 (using `continue`):**
```go
for i := 1; i <= 100; i++ {
	if i%7 != 0 {
		continue
	}
	fmt.Println(i)
}
```
- If the number is **not divisible by 7**, `continue` skips to the next iteration.
- Only numbers divisible by `7` are printed.

### ✅ **Output:**
```
7
14
21
28
...
98
```

---

## **3. Stop the loop when the number reaches 10 (using `break`):**
```go
for i := 1; i <= 100; i++ {
	if i == 10 {
		break
	}
	fmt.Println(i)
}
```
- The loop runs from `1` to `100`.
- When `i == 10`, the `break` statement **exits the loop immediately**.

### ✅ **Output:**
```
1
2
3
...
9
```
Then, it prints:

```
After loop
```

---

## **4. Iterating through a slice using index-based for loop:**
```go
nums := []int{1, 2, 3, 4, 5}
for i := 0; i < len(nums); i++ {
	fmt.Println(nums[i])
}
```
- A traditional loop using an index (`i`) to access each element in the slice `nums`.
- `len(nums)` returns the length of the slice.

### ✅ **Output:**
```
1
2
3
4
5
```

---

## **5. Using `range` to loop with both index and value:**
```go
for index, element := range nums {
	fmt.Printf("Index: %d Element: %d\n", index, element)
}
```
- `range` allows iterating over the slice while accessing both the **index** and **value**.
- The `index` holds the position, and `element` holds the value.

### ✅ **Output:**
```
Index: 0 Element: 1
Index: 1 Element: 2
Index: 2 Element: 3
Index: 3 Element: 4
Index: 4 Element: 5
```

---

## **6. Using `range` to loop with only the value:**
```go
for _, element := range nums {
	fmt.Printf("Element: %d\n", element)
}
```
- The underscore (`_`) is used to **ignore the index**.
- Only the value `element` is printed.

### ✅ **Output:**
```
Element: 1
Element: 2
Element: 3
Element: 4
Element: 5
```

---

## **7. Nested loop to iterate over a matrix (2D slice):**
```go
matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
for _, row := range matrix {
	for _, column := range row {
		fmt.Printf("%d ", column)
	}
	fmt.Println()
}
```
- The outer loop iterates through each **row**.
- The inner loop iterates through each **column** in the current row.
- The `fmt.Printf` statement prints numbers in the same line, and `fmt.Println()` moves to the next line after each row.

### ✅ **Output:**
```
1 2 3 
4 5 6 
7 8 9
```

---

## 🎯 **Summary of Key Concepts:**
| Concept            | Explanation |
|----------------|-----------------------------|
| Basic `for` loop | Standard loop with initialization, condition, and increment. |
| `continue` | Skips the current iteration and moves to the next. |
| `break` | Exits the loop immediately. |
| Index-based loop | Accesses elements using `nums[i]`. |
| `range` loop | Simplifies iteration through slices and maps. |
| Nested loop | Iterates through a matrix (2D slice). |

---

Let me know if you'd like additional examples or improvements! 😊