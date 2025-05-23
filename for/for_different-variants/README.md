This Go program demonstrates **different types of `for` loops** in Go, which is the only loop construct in the language. It covers **basic iteration, conditional loops, infinite loops, `continue` and `break` statements, and iterating over arrays.**

---

## ✅ **Step-by-step Explanation:**

---

## 🎯 **1. Basic `for` loop with condition:**
```go
i := 1

for i <= 3 {
	fmt.Println(i)
	i++ // Equivalent to i = i + 1
}
```
- This loop **continues as long as `i` is less than or equal to 3**.
- It prints `i`, then increments it by 1.

### 🔹 Output:
```
1
2
3
```

---

## 🎯 **2. `for` loop with initialization, condition, and increment:**
```go
for j := 7; j <= 9; j++ {
	fmt.Println(j)
}
```
- This loop **initializes `j` to 7**, runs as long as `j` is **less than or equal to 9**, and **increments `j` by 1** after each iteration.

### 🔹 Output:
```
7
8
9
```

---

## 🎯 **3. Infinite `for` loop with `break`:**
```go
for {
	fmt.Println("loop")
	break
}
```
- This creates an **infinite loop**.
- `break` is used to **exit the loop immediately.**

### 🔹 Output:
```
loop
```

---

## 🎯 **4. `for` loop with `continue` (skipping even numbers):**
```go
for n := 0; n <= 5; n++ {
	if n%2 == 0 {
		continue
	}
	fmt.Println(n)
}
```
- `n%2 == 0` checks if `n` is even.
- `continue` **skips the current iteration** when `n` is even, so only **odd numbers are printed.**

### 🔹 Output:
```
1
3
5
```

---

## 🎯 **5. Iterating over an array with `range` (index and value):**
```go
var fruits = [3]string{"apple", "pear", "baban"}
for index, value := range fruits {
	fmt.Println(index, value)
}
```
- `range` returns both **index** and **value** from the `fruits` array.

### 🔹 Output:
```
0 apple
1 pear
2 baban
```

---

## 🎯 **6. Iterating over an array with `range` (only values):**
```go
for _, value := range fruits {
	fmt.Println(value)
}
```
- The underscore (`_`) is used to **ignore the index**.
- Only **the values are printed.**

### 🔹 Output:
```
apple
pear
baban
```

---

## 🎯 **7. Classic `for` loop using `len()`:**
```go
for i := 0; i < len(fruits); i++ {
	fmt.Println(fruits[i])
}
```
- This is a **traditional `for` loop** that **manually accesses each element** using the index `i`.

### 🔹 Output:
```
apple
pear
baban
```

---

## ✅ **Final Output of the Program:**
```
1
2
3
7
8
9
loop
1
3
5
0 apple
1 pear
2 baban
apple
pear
baban
apple
pear
baban
```

---

## 🔥 **Key Concepts:**
| Loop Type             | Description                                | Control Statements |
|----------------|---------------------------------|------------------|
| `for` with condition | Runs until the condition is false | `i++` |
| Classic `for` loop | Initializes, checks condition, and increments | `j++` |
| Infinite `for` loop | Runs forever unless broken | `break` |
| `for` with `continue` | Skips specific iterations | `continue` |
| `range` loop | Iterates over arrays/slices/maps | `range` |

---

Let me know if you'd like more examples or modifications! 😊