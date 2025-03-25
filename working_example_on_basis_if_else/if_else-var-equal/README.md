Let's break down the Go code step by step and explain the rules.

### Code:

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	c := a + b
	if c == 10 {
		fmt.Println("==10")
	} else {
		fmt.Println("!=10")
	}
}
```

### **Explanation:**

1. **Variable Declarations:**
   ```go
   var a, b int
   ```
   - Here, two integer variables `a` and `b` are declared.

2. **Reading User Input:**
   ```go
   fmt.Scanf("%d", &a)
   fmt.Scanf("%d", &b)
   ```
   - `fmt.Scanf` is used to read input from the user. 
   - The format specifier `"%d"` tells the program to read an integer.
   - The first `fmt.Scanf` reads an integer input for `a`, and the second reads an integer input for `b`.

   **In the example**, the user inputs:
   ```
   5
   6
   ```
   So, `a = 5` and `b = 6`.

3. **Adding the Two Numbers:**
   ```go
   c := a + b
   ```
   - This line calculates the sum of `a` and `b` and stores the result in `c`.
   - In this case:
     - `a = 5` and `b = 6`, so `c = 5 + 6 = 11`.

4. **If-Else Statement:**
   ```go
   if c == 10 {
       fmt.Println("==10")
   } else {
       fmt.Println("!=10")
   }
   ```
   - The program checks if `c` is equal to `10`.
   - The condition `if c == 10` checks if the value of `c` is 10.
   - In our case, since `c = 11`, the condition `c == 10` is **false**.
   - As a result, the program will execute the `else` block, which prints `"!=10"`.

### **Output:**

- Since the sum `c` is not equal to `10`, the output will be:
  ```
  !=10
  ```

### **Summary of Execution:**

- **Input:** 
  - The user enters `5` and `6` for `a` and `b`.
- **Processing:**
  - `c` is calculated as `5 + 6 = 11`.
  - The program then checks if `c == 10`. Since `c` is `11`, this condition is false.
  - Therefore, the program executes the `else` block.
- **Output:** 
  - The program prints `!=10`.

### **Key Concepts:**

1. **Input Handling:**
   - `fmt.Scanf` reads input from the user. The values are stored in the variables `a` and `b`.
   
2. **Conditionals (`if-else`):**
   - The `if` statement checks whether `c` is equal to `10`. If true, it prints `"==10"`. Otherwise, it prints `"!=10"`.
   
3. **Basic Arithmetic Operations:**
   - The addition operation (`a + b`) computes the sum of the two numbers, and the result is stored in the variable `c`.

### **Conclusion:**
- The program reads two integer inputs, computes their sum, and checks whether the sum is equal to `10`. If it is, it prints `"==10"`, otherwise it prints `"!=10"`. Since the sum in this case is `11`, the output is `"!=10"`.