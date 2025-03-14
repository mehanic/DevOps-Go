In this Go code, you're demonstrating various **constant types** and how they work when declared. Let's break down the rules and explain what's happening:

### Code Breakdown:

```go
package main

import "fmt"

func main() {

	const x = 3
	const y = 5.6
	const z = true
	const t = "test"

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)

}
```

### Explanation of the Rules:

1. **Constant Declaration:**
   ```go
   const x = 3
   const y = 5.6
   const z = true
   const t = "test"
   ```
   - Here, four **constants** are being declared:
     - `x = 3`: An **integer** constant with the value `3`. The Go compiler automatically infers that `x` has the type `int`.
     - `y = 5.6`: A **floating-point** constant with the value `5.6`. The Go compiler infers the type `float64` because floating-point literals in Go default to `float64`.
     - `z = true`: A **boolean** constant with the value `true`. The type inferred is `bool`.
     - `t = "test"`: A **string** constant with the value `"test"`. The inferred type is `string`.

2. **Constant Types Are Inferred by the Compiler:**
   - Go **infers the types** of constants based on their values. In this case:
     - `x` is an integer, so Go infers `x` as `int`.
     - `y` is a floating-point number, so Go infers `y` as `float64`.
     - `z` is a boolean literal, so Go infers `z` as `bool`.
     - `t` is a string literal, so Go infers `t` as `string`.
   - **Important:** Constants in Go can either have an explicit type (like `const x int = 5`) or the type can be inferred, as is done here.

3. **Printing Constants:**
   ```go
   fmt.Printf("%T, %v\n", x, x)
   fmt.Printf("%T, %v\n", y, y)
   fmt.Printf("%T, %v\n", z, z)
   fmt.Printf("%T, %v\n", t, t)
   ```
   - `fmt.Printf("%T, %v\n", x, x)` will print:
     - `%T`: The **type** of `x`, which is `int`.
     - `%v`: The **value** of `x`, which is `3`.
   - Similarly, it prints the type and value for the other constants (`y`, `z`, `t`):
     - For `y`: Type `float64` and value `5.6`.
     - For `z`: Type `bool` and value `true`.
     - For `t`: Type `string` and value `"test"`.

4. **Output:**
   The output of the program will be:
   ```
   int, 3
   float64, 5.6
   bool, true
   string, test
   ```

### Summary of Rules:

1. **Constant Declaration:** Constants can be declared with or without an explicit type. If no type is given, Go infers the type from the value.
   
2. **Constant Types Are Inferred:**
   - When you assign a value to a constant, Go will infer the type based on the value's literal type:
     - Integer values default to `int`.
     - Floating-point values default to `float64`.
     - Boolean values default to `bool`.
     - String literals default to `string`.

3. **Printing Constants:**
   - The `fmt.Printf("%T, %v\n", x, x)` statement is used to print the **type** (`%T`) and **value** (`%v`) of a constant.
   
4. **Types Are Immutable:** Once a constant is declared, its value and type are fixed throughout the program.

### Key Takeaways:

- **Constants are typed either implicitly or explicitly**.
- **Go's type inference system** automatically detects the correct type for constants based on their values.
- You can use `fmt.Printf` with the `%T` format to print the type of a constant and `%v` to print its value.
