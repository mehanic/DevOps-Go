In this Go code, you are using **constants** to store values that remain unchanged throughout the execution of the program. Let's break down the rules and explain what's happening here.

### Code Breakdown:

```go
package main

import "fmt"

const (
    pi       = 3.14
    language = "Go"
)

func main() {
    fmt.Println(pi)
    fmt.Println(language)
}
```

### Explanation of the Rules:

1. **Constant Declaration:**
   ```go
   const (
       pi       = 3.14
       language = "Go"
   )
   ```
   - In this code, two **constants** are defined within a `const` block:
     - `pi = 3.14`: The constant `pi` is assigned the value `3.14`. Since `3.14` is a floating-point number, the **type of `pi`** is inferred by the Go compiler to be `float64`.
     - `language = "Go"`: The constant `language` is assigned the string value `"Go"`. The **type of `language`** is inferred to be `string`.
   - **Multiple constants** can be declared in a single `const` block, separated by line breaks or commas. The constants are grouped together for clarity and simplicity.

2. **Untyped Constants:**
   - Both constants `pi` and `language` are **untyped constants** because you did not explicitly define their types. Go will **implicitly infer** their types based on the assigned values:
     - `pi` will be inferred as a `float64` because `3.14` is a floating-point number.
     - `language` will be inferred as a `string` because `"Go"` is a string literal.
   - Go's type inference system automatically handles this, so you don't need to explicitly declare types unless necessary.

3. **Constants are Immutable:**
   - Constants in Go are **immutable**. Once a value is assigned to a constant, it cannot be changed during the program's execution.
   - For example, if you tried to assign a new value to `pi` or `language` anywhere in the code, Go would raise a compile-time error.

4. **Constants are Evaluated at Compile Time:**
   - Constants in Go are **evaluated at compile time**, meaning their values are known and fixed before the program starts running. This is why you can use constants in places that require constant values, like array sizes, case labels in a `switch`, etc.

5. **Printing Constants:**
   - In the `main` function, you print the values of `pi` and `language` using `fmt.Println()`:
     - `fmt.Println(pi)` will output the value of `pi` (which is `3.14`).
     - `fmt.Println(language)` will output the value of `language` (which is `"Go"`).

6. **Output:**
   - The output of this program will be:
     ```
     3.14
     Go
     ```

### Summary of Rules:

- Constants are **immutable values** that are assigned at compile time.
- They can be **implicitly typed** based on the value assigned (e.g., `float64` for `3.14` and `string` for `"Go"`).
- Constants cannot be **modified** during the program execution.
- Constants are typically used for values that remain constant throughout the lifetime of the program, such as mathematical constants, configuration values, or fixed strings.
- The **`fmt.Println()`** function is used to print the constant values to the console.

In conclusion, this code defines two constants `pi` and `language`, prints them, and demonstrates the use of constants in Go. The constants are immutable, evaluated at compile-time, and inferred by the compiler based on their assigned values.