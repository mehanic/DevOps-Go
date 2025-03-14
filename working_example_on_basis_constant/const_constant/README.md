In Go, constants have specific rules that govern how they can be used, and one of the key rules is that **constants cannot be reassigned once they are declared**. Let's break this down based on the code you've provided:

### Code Breakdown:

```go
package main

import "fmt"

func main() {
	const Pi = 3.14
	// Pi = 400 // enbling this line will throw an error: cannot assign to Pi
	fmt.Println("Pi=", Pi)
}
```

### Explanation:

1. **Constant Declaration:**
   ```go
   const Pi = 3.14
   ```
   - Here, we declare a constant named `Pi` with the value `3.14`. In Go, when you use the `const` keyword, you're defining a **constant** that holds a value that cannot be changed during the program's execution.
   - Go **automatically infers** the type of the constant based on the value. In this case, `Pi` will have the type `float64` because the literal `3.14` is a floating-point number, and Go defaults floating-point literals to `float64`.

2. **Attempt to Modify a Constant (Commented Out Code):**
   ```go
   // Pi = 400 // enabling this line will throw an error: cannot assign to Pi
   ```
   - This line attempts to **reassign** the value of `Pi` to `400`. However, this results in an error if you uncomment the line because **constants in Go are immutable**. Once a constant is set to a value, it **cannot** be reassigned.
   - **Why does this cause an error?**
     - Constants in Go are treated as **immutable** values. You cannot change their value once they are defined. This is different from variables, which can be reassigned during the program's execution.
     - The error message you would see if this line is uncommented is: `cannot assign to Pi`, which explicitly tells you that constants cannot be modified after their initial assignment.

3. **Printing the Constant:**
   ```go
   fmt.Println("Pi=", Pi)
   ```
   - This line prints the value of the constant `Pi` to the console. Since `Pi` is a constant, its value remains `3.14` throughout the program.

### Key Rule: **Immutability of Constants**

- Once you declare a constant with the `const` keyword, its value **cannot be changed**. The constant’s value is **set at compile-time** and remains the same throughout the program’s lifetime.
  
  In contrast to **variables**, which can be assigned new values at any time during the program’s execution, constants are **fixed** and cannot be modified.

### Why Constants Are Immutable:

1. **Efficiency**: Constants are often used for values that are meant to remain constant throughout the program, such as mathematical constants (`Pi`, `e`, etc.), fixed configuration values, or flags. Making them immutable ensures that they cannot accidentally be modified, preventing bugs or unintended behavior.
   
2. **Compile-Time Evaluation**: Constants in Go are evaluated at compile-time, not runtime. This makes them more efficient, as their values are resolved before the program is run. Allowing reassignment could break this compile-time evaluation.

### Summary of Rules:

1. **Constants are immutable**: Once a constant is assigned a value, it **cannot be changed** during the execution of the program.
2. **Type inference**: Constants automatically infer their type from the assigned value (e.g., `3.14` becomes a `float64`).
3. **Attempting to reassign a constant** results in a compile-time error (`cannot assign to Pi`).

### Output of the Code:

When you run this code (with the reassignment line commented out), the output will be:
```
Pi= 3.14
```

This shows the constant value of `Pi`, which remains the same throughout the execution of the program.