In the Go code you provided, there are two constants defined at the beginning of the program and then printed out in the `main` function. Let's break down the key concepts and rules used here:

### Code Breakdown:

```go
const (
    aa = "hello"
    b  = 44
)

func main() {
    println(aa)
    println(b)
}
```

1. **Constant Declaration:**
   ```go
   const (
       aa = "hello"
       b  = 44
   )
   ```
   - In this code, **constants** are declared within a `const` block.
     - `aa = "hello"`: Here, `aa` is a constant that holds the string value `"hello"`. The **type of `aa`** is implicitly set by the compiler to `string` because the value is a string.
     - `b = 44`: Here, `b` is a constant that holds the integer value `44`. The **type of `b`** is implicitly set to `int` by the compiler because the value is an integer.

2. **Constant Values:**
   - Constants in Go are **immutable** (they cannot be changed once they are set).
   - They are evaluated **at compile-time**, meaning that their values are fixed when the program is compiled.

3. **Constant Types:**
   - **Untyped Constants:** Both `aa` and `b` are **untyped constants** because they do not have explicit types assigned to them. In Go, **untyped constants** can be inferred by the Go compiler based on their values. 
     - `aa` is inferred to be of type `string` because it is assigned a string value.
     - `b` is inferred to be of type `int` because it is assigned an integer value.

4. **Usage of Constants:**
   - In the `main` function, the values of these constants are printed using `println(aa)` and `println(b)`.
     - `println(aa)` will print the string `"hello"`.
     - `println(b)` will print the integer `44`.

5. **Implicit Type Inference:**
   - **No need to specify types explicitly** when declaring constants if the value type is clear. In this case:
     - `aa` is automatically treated as a `string`.
     - `b` is automatically treated as an `int`.

### Key Rules:
- **Constants are fixed values** that are assigned at compile time and **cannot be modified during runtime**.
- Constants can be declared **without specifying types** if their values are obvious (like `"hello"` being a `string` or `44` being an `int`).
- Go **infers the type** of constants based on their values.
- The **`println` function** is used to print the values to the console. Note that `println` outputs the values without formatting. If you need formatted output, you would use `fmt.Println` instead.

### Output of the Program:
```go
hello
44
```

- `"hello"` will be printed because `aa` is a constant of type `string` with the value `"hello"`.
- `44` will be printed because `b` is a constant of type `int` with the value `44`.