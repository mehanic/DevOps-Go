The Go program you provided demonstrates the use of **iota**, which is a special identifier used in Go to generate a sequence of constant values. Here's a breakdown of the code and the rules behind how **iota** works:

### Code Explanation:

```go
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)
```

1. **`const` Block:**
   - A `const` block is used to group together constant values. In Go, the `const` keyword is used to define constants, which are values that cannot change during the execution of the program.

2. **`iota`:**
   - `iota` is a special keyword in Go that is used inside constant declarations. It starts at `0` and automatically increments by 1 with each constant in the block.
   
   - When `iota` is used in a `const` block:
     - The first constant (`a`) gets the value `0` (the initial value of `iota`).
     - The second constant (`b`) gets the value `1` because `iota` increments after each constant.
     - The third constant (`c`) gets the value `2`, and so on.
   
   In the code you provided, `iota` is used with every constant definition, which increments its value for each line within the same `const` block.

### Key Details:

1. **Automatic Increment:**
   - The main feature of `iota` is that it automatically increments the value for each constant in the same `const` block. You don’t need to manually assign a value to each constant; `iota` will handle the incrementation for you.
   - This behavior makes it very useful when you want a sequence of constant values.

2. **`iota` Example without Repeating:**
   The following code shows an alternate way of defining constants without repeating `iota` explicitly:

   ```go
   const (
       a = iota // 0
       b        // 1
       c        // 2
   )
   ```

   In this case, you don't need to repeat `iota` for each line. By default, `iota` will increment by 1 for each line in the `const` block.

3. **Output:**
   The output of the program will be:
   ```
   0
   1
   2
   ```
   This is because:
   - `a = iota` assigns `a` the value `0` (since `iota` starts at `0`).
   - `b = iota` assigns `b` the value `1` (since `iota` increments).
   - `c = iota` assigns `c` the value `2` (since `iota` increments again).

### **Practical Use of `iota`:**

In Go, `iota` is often used to create **enumerations** (similar to `enum` in other languages like C++). It allows you to create sequences of related constants, which is useful when defining flags, error codes, or any other scenario where you need a series of sequential values.

For example:

```go
const (
    Red = iota // 0
    Green      // 1
    Blue       // 2
)
```

Here, the constants `Red`, `Green`, and `Blue` represent different color values, with `iota` automatically assigning them the values `0`, `1`, and `2` respectively.

### Conclusion:
- **iota** is a powerful tool in Go to automatically generate sequential constant values.
- It simplifies the process of creating a series of related constants, like in enumerations.
- Each time `iota` is used in a `const` block, it increments by 1, starting from 0.
- By omitting `iota` on subsequent constants within the same block, Go automatically increments the value for you.