Let's go through each of the Go programs and explain the rules and concepts:

---

### **First Program Explanation:**
```go
package main

import "fmt"

func main() {
	// Arifmetik Operator: +, -, /, *, %
	var (
		a int = 10
		b float32 = 22.23
		c float64 = 23.234
	)

	fmt.Println(a + int(b))
	fmt.Println(b + float32(c))

	var age uint8 = 255
	fmt.Println(age)

	// int8, int16, int32, int64, int
	// uint8, uint16, uint32, uint64, uint
}
```

#### Explanation:

1. **Variable Declaration:**
   - `a` is an integer (`int`), `b` is a `float32`, and `c` is a `float64`.
   - Go supports both signed (`int`, `int8`, `int16`, `int32`, `int64`) and unsigned (`uint`, `uint8`, `uint16`, `uint32`, `uint64`) integers. Signed integers can hold both positive and negative numbers, while unsigned integers can only hold positive numbers and zero.

2. **Arithmetic Operations:**
   - `a + int(b)`: The `b` variable, a `float32`, is explicitly converted to an `int` before performing the addition. In Go, type conversions must be explicit, unlike some other languages where implicit conversion might happen. This results in truncation of the decimal part of `b` when it is cast to `int`.
   - `b + float32(c)`: Here, `c` is a `float64` and is explicitly converted to `float32` to perform the addition. Type conversions are required when dealing with different numeric types (e.g., `float32` and `float64`).

3. **Unsigned Integers:**
   - `age uint8 = 255`: Here, the variable `age` is an unsigned 8-bit integer (`uint8`), and its value is set to 255, which is the maximum value `uint8` can hold (since `uint8` can only store values from 0 to 255).
   
---

### **Second Program Explanation:**
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(math.Floor(3.5))
}
```

#### Explanation:

1. **String and Mathematical Operation:**
   - The program prints "hello world" as a simple string output.
   - `math.Floor(3.5)` rounds down the number `3.5` to the nearest integer. The `math.Floor` function returns the largest integer less than or equal to the given value, which in this case is `3`.

   **Output:**
   - "hello world"
   - 3.0 (since the `Floor` function returns a `float64`, even though it's an integer value, it will be printed as `3.0`).

---

### **Third Program Explanation:**
```go
package main

import "fmt"

func main() {

	// var num int = 10
	// // var a int

	// // num = "ten"

	// fmt.Println(num)
	// fmt.Printf("%T\n", num)

	// int -> deafult: 0
	var a, b, c int = 10, 12, 14
	fmt.Println(a, b, c)

	// float32 va float64 -> deafult: 0
	var (
		d float32 = 12.24
		e float64 = 54.4
	)
	fmt.Println("Float->", d, e)

	// boolean -> deafult: false
	var isMarried bool = true

	fmt.Println("Boolean->", isMarried)
	// string -> deafult: ""
	var username string = "admin"
	fmt.Println("String->", username)
}
```

#### Explanation:

1. **Variable Initialization and Types:**
   - Go has a strong typing system, which means variables have specific types, and you cannot assign a value of one type to a variable of another type without explicit conversion.
   - The program declares variables `a`, `b`, and `c` as `int` and initializes them with `10`, `12`, and `14`, respectively.
   
2. **Default Values:**
   - **Integer:** `int` variables have a default value of `0` if not explicitly initialized.
   - **Float:** `float32` and `float64` types have a default value of `0` when not initialized.
   - **Boolean:** The default value of `bool` is `false`.
   - **String:** The default value of a `string` is an empty string (`""`).

3. **Output of Types and Values:**
   - The program prints the values and types of the variables. When variables are initialized, they print the values as expected.
   - The types and values are printed using the `fmt.Println` and `fmt.Printf` functions.

   **Output:**
   - `10 12 14`
   - `Float-> 12.24 54.4`
   - `Boolean-> true`
   - `String-> admin`

### Key Takeaways:
1. **Type Safety:** Go is statically typed, meaning the type of a variable is known at compile-time and does not change at runtime. This helps prevent errors but requires explicit type conversions when needed.
   
2. **Default Values:** 
   - **Int** defaults to `0`.
   - **Float32/Float64** defaults to `0`.
   - **Boolean** defaults to `false`.
   - **String** defaults to an empty string (`""`).

3. **Explicit Type Conversion:** Go requires explicit type conversions, unlike languages that allow implicit conversions. This ensures clarity and avoids unexpected behavior.

---

These rules and principles are fundamental for working with types and variables in Go, and they help ensure type safety and correct behavior in Go programs.