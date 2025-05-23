Let's break down and explain the code you've shared, which involves several control statements in Go, including comparison operations, `if`, `else`, `else if`, `switch`, and `for` loops.

### **Code Breakdown:**

```go
package main

import "fmt"

/*
control statements
if /else
if/elseif /else
switch statement
*/

func main() {

	a := 4
	b := 5

	// Comparison
	fmt.Printf("a == b, %v\n", a == b) // false
	fmt.Printf("a == a, %v\n", a == a) // true
	fmt.Printf("a != b, %v\n", a != b) // true
	fmt.Printf("a > b, %v\n", a > b)   // false
	fmt.Printf("a < b, %v\n", a < b)   // true
	fmt.Printf("a <= b, %v\n", a <= b) // true
	fmt.Printf("a <= b, %v\n", a <= a) // true
	fmt.Printf("a >= b, %v\n", a >= a) // true

	fmt.Println("--------------------")

	// if without else statement
	if a > b {
		fmt.Printf("if statement: a is greater than b , %v\n", a > b)
	}
	fmt.Printf("else statement: a is lesser than b , %v\n", a > b)

	// if - else if - else statement
	if a > b {
		fmt.Printf("if: a is greater than b , %v\n", a > b)
	} else if a == b {
		fmt.Printf("if else: a is same as  b , %v\n", a > b)
	} else {
		fmt.Printf("else: a is lesser than b , %v\n", a > b)
	}
	fmt.Println("--------------------")

	//switch case
	switch {
	case a < 5:
		fmt.Println("switch: a is less than 5")
	case a == 5:
		fmt.Println("switch: a is equal to 5")
	default:
		fmt.Println("switch: a is greater than 5")
	}

	switch a {
	case 5:
		fmt.Println("switch: a is equal to 5")
	}

	// for loop
	for a < 10 {
		fmt.Printf("For: a - %v\n", a)
		a++
	}

	for i := 0; i < b; i++ {
		fmt.Printf("for loop b: %v\n", i)
	}

}
```

### **Explanation of Each Section:**

#### 1. **Comparison Operations:**

```go
fmt.Printf("a == b, %v\n", a == b) // false
fmt.Printf("a == a, %v\n", a == a) // true
fmt.Printf("a != b, %v\n", a != b) // true
fmt.Printf("a > b, %v\n", a > b)   // false
fmt.Printf("a < b, %v\n", a < b)   // true
fmt.Printf("a <= b, %v\n", a <= b) // true
fmt.Printf("a <= b, %v\n", a <= a) // true
fmt.Printf("a >= b, %v\n", a >= a) // true
```
- **Comparison operators** are used to compare two values and return a boolean result:
  - `==`: Checks if two values are equal.
  - `!=`: Checks if two values are not equal.
  - `>`: Checks if the left value is greater than the right.
  - `<`: Checks if the left value is less than the right.
  - `<=`: Checks if the left value is less than or equal to the right.
  - `>=`: Checks if the left value is greater than or equal to the right.

Here, the values of `a` and `b` are `4` and `5`, respectively, so the comparisons are:
- `a == b` returns `false`.
- `a == a` returns `true`.
- `a != b` returns `true`.
- `a > b` returns `false`.
- `a < b` returns `true`.
- `a <= b` returns `true`.
- `a <= a` returns `true`.
- `a >= a` returns `true`.

#### 2. **If Without Else Statement:**
```go
if a > b {
	fmt.Printf("if statement: a is greater than b , %v\n", a > b)
}
fmt.Printf("else statement: a is lesser than b , %v\n", a > b)
```
- **Explanation:**
  - The `if` statement checks whether `a > b`. Since `a` is `4` and `b` is `5`, the condition is false, and the `if` block is **not executed**.
  - The second `fmt.Printf` prints `"else statement: a is lesser than b , false"`, as the `if` condition was false.

#### 3. **If - Else If - Else Statement:**
```go
if a > b {
	fmt.Printf("if: a is greater than b , %v\n", a > b)
} else if a == b {
	fmt.Printf("if else: a is same as  b , %v\n", a > b)
} else {
	fmt.Printf("else: a is lesser than b , %v\n", a > b)
}
```
- **Explanation:**
  - The `if` condition `a > b` is false.
  - The `else if` condition `a == b` is also false.
  - The `else` block is executed because all the previous conditions failed. The statement `"else: a is lesser than b , false"` is printed, since the condition `a > b` was false.

#### 4. **Switch Statement:**
```go
switch {
case a < 5:
	fmt.Println("switch: a is less than 5")
case a == 5:
	fmt.Println("switch: a is equal to 5")
default:
	fmt.Println("switch: a is greater than 5")
}
```
- **Explanation:**
  - This `switch` statement checks conditions without using a specific variable. It checks the conditions sequentially.
  - The first condition (`a < 5`) is true, so the statement `"switch: a is less than 5"` is printed.
  - The other cases are skipped because the first condition was satisfied.

#### 5. **Another Switch on Variable:**
```go
switch a {
case 5:
	fmt.Println("switch: a is equal to 5")
}
```
- **Explanation:**
  - This switch checks the value of `a` directly.
  - Since `a` is `4`, no case matches, so nothing is printed.

#### 6. **For Loop (First Example):**
```go
for a < 10 {
	fmt.Printf("For: a - %v\n", a)
	a++
}
```
- **Explanation:**
  - This is a **basic `for` loop**. It prints the value of `a` and then increments `a` by 1 in each iteration.
  - The loop continues as long as `a < 10`. It prints:
    ```
    For: a - 4
    For: a - 5
    For: a - 6
    For: a - 7
    For: a - 8
    For: a - 9
    ```
  - After the loop, `a` becomes `10`, and the loop stops.

#### 7. **For Loop (Second Example):**
```go
for i := 0; i < b; i++ {
	fmt.Printf("for loop b: %v\n", i)
}
```
- **Explanation:**
  - This is a **`for` loop with an index variable (`i`)**.
  - The loop runs from `i = 0` to `i = 4` (since `b = 5`), printing:
    ```
    for loop b: 0
    for loop b: 1
    for loop b: 2
    for loop b: 3
    for loop b: 4
    ```

### **Summary of Control Statements:**
- **Comparison operators** are used to compare values and return a boolean.
- **`if` statements** execute code based on a condition.
- **`else if` statements** allow multiple conditions to be checked in sequence.
- **`else` statements** execute code when no previous conditions are true.
- **`switch` statements** evaluate a value or conditions and execute code based on matched cases.
- **`for` loops** repeat code until a condition is no longer true.

### **Output:**
```
a == b, false
a == a, true
a != b, true
a > b, false
a < b, true
a <= b, true
a <= b, true
a >= b, true
--------------------
else statement: a is lesser than b , false
else: a is lesser than b , false
--------------------
switch: a is less than 5
For: a - 4
For: a - 5
For: a - 6
For: a - 7
For: a - 8
For: a - 9
for loop b: 0
for loop b: 1
for loop b: 2
for loop b: 3
for loop b: 4
```