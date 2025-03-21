This Go program demonstrates several key concepts related to printing, variable scope, naming conventions, and shadowing. Let's break it down step by step:

---

## 1. **Print Functions in Go:**

### `fmt.Print()`
- Prints text on the same line without adding a new line at the end.
- Example: 
  ```go
  fmt.Print("Hello World")
  ```
  Output: `Hello World`

---

### `fmt.Println()`
- Prints text and automatically adds a new line at the end.
- Example:
  ```go
  fmt.Println("Hello World")
  ```
  Output:
  ```
  Hello World
  ```

---

### `fmt.Printf()`
- Allows **formatted output**, similar to `printf` in C.
- You can use **format specifiers** like `%v`, `%T`, `%d`, etc.
- Example:
  ```go
  fmt.Printf("My name is: %v", name)
  ```
  Output:
  ```
  My name is: Deniz
  ```

---

### Why this output is wrong?
```go
fmt.Printf("My name is:", name)
```
- The `%v` placeholder is missing, so Go doesn’t know where to put the `name` variable.
- The output is:
  ```
  My name is: %!(EXTRA string=Deniz)
  ```

---

## 2. **Using Format Specifiers:**
- `%v`: Prints the **value**.
- `%T`: Prints the **type** of the variable.
- `%d`: Prints an **integer** in decimal format.
- `%c`: Prints the **character** corresponding to the ASCII/Unicode value.

---

## 3. **Scope and Shadowing:**
- **Scope** refers to the visibility of variables in different code blocks.
- **Shadowing** occurs when a variable declared in an **inner block** has the same name as a variable in an **outer block**, temporarily hiding the outer variable.

### Example:
```go
x := 5

{
	x := 10
	fmt.Println(x) // Outputs 10 (the inner x)
}

fmt.Println(x) // Outputs 5 (the outer x)
```
- The inner `x` **shadows** the outer `x` within the block, but the outer `x` remains unchanged outside the block.

---

## 4. **Naming Conventions:**
- Use **CamelCase** (e.g., `myVariable`).
- Abbreviations like `URL`, `HTTP` should be **uppercase** (e.g., `myURL`, `myHTTP`).
- Short variable names like `i`, `j`, `k` are common for loop counters.
- Avoid non-English characters for better compatibility (e.g., Turkish characters).

---

## ✅ **Final Output:**
```
Hello WorldHello WorldHello World
Deniz
Deniz
DenizMy name is:Deniz
My name is: Deniz
My name is: %!(EXTRA string=Deniz)
My name is: DenizMy name is: stringMy name is: Deniz string
10
5
```

---

Let me know if you'd like more examples or deeper explanations! 😊