Let's break down the Go code and explain the rules step by step:

### Code:

```go
package main

import "fmt"

func main() {
	a := 10
	b := 22
	c := 5
	if a > b && a > c {
		fmt.Println(a)
	}
	if b > a && b > c {
		fmt.Println(b)
	}
	if c > a && c > b {
		fmt.Println(c)
	}
}
```

### **Explanation:**

1. **Variable Initialization:**
   ```go
   a := 10
   b := 22
   c := 5
   ```
   - `a` is assigned the value `10`.
   - `b` is assigned the value `22`.
   - `c` is assigned the value `5`.

2. **First `if` condition:**
   ```go
   if a > b && a > c {
       fmt.Println(a)
   }
   ```
   - This condition checks if `a` is greater than both `b` and `c`. In other words, it checks if `a > b` **and** `a > c`.
   - The condition evaluates to:
     - `a > b` → `10 > 22` is **false**.
     - `a > c` → `10 > 5` is **true**.
   - Since the condition is a logical **AND** (`&&`), both parts must be true for the overall condition to be true. Since `a > b` is false, the entire condition evaluates to **false**.
   - Because the condition is false, the `fmt.Println(a)` inside the block will **not** execute.

3. **Second `if` condition:**
   ```go
   if b > a && b > c {
       fmt.Println(b)
   }
   ```
   - This condition checks if `b` is greater than both `a` and `c`. In other words, it checks if `b > a` **and** `b > c`.
   - The condition evaluates to:
     - `b > a` → `22 > 10` is **true**.
     - `b > c` → `22 > 5` is **true**.
   - Since both parts are true, the overall condition evaluates to **true**.
   - Because the condition is true, `fmt.Println(b)` inside the block will execute and print the value of `b`, which is `22`.

4. **Third `if` condition:**
   ```go
   if c > a && c > b {
       fmt.Println(c)
   }
   ```
   - This condition checks if `c` is greater than both `a` and `b`. In other words, it checks if `c > a` **and** `c > b`.
   - The condition evaluates to:
     - `c > a` → `5 > 10` is **false**.
     - `c > b` → `5 > 22` is **false**.
   - Since both parts of the condition are false, the overall condition evaluates to **false**.
   - Because the condition is false, `fmt.Println(c)` inside the block will **not** execute.

### **Summary of Execution Flow:**
1. **First `if` condition** is **false**, so nothing is printed.
2. **Second `if` condition** is **true**, so `b` (which is `22`) is printed.
3. **Third `if` condition** is **false**, so nothing is printed.

### **Output:**
```
22
```

### **Important Concepts:**

1. **Logical AND (`&&`):**
   - The `&&` operator is used to check if both conditions are **true**. If both parts of the condition are true, the entire condition is true. If either part is false, the entire condition becomes false.
   - In the first and third `if` statements, one part of the condition is false, so the entire condition becomes false.

2. **Evaluating Conditions:**
   - Go evaluates conditions in a left-to-right manner. If the first condition in an `&&` expression is false, the second condition is not evaluated because the entire condition can never be true.

3. **Conditional Execution:**
   - Only the `if` block whose condition evaluates to **true** will execute. In this case, only the second `if` block prints anything.

### **Conclusion:**
- The program evaluates each condition one by one.
- The second `if` block is the only one where both conditions are true, so it prints `22`.
