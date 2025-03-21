This Go program demonstrates **variable declaration**, **type inference**, and **type checking**. Let's break it down step by step:

---

## ✅ **1. Variable Declaration:**
```go
var x = 10
var name = "kira"
var z string = "random string"
var y = `some string, "random string"`
```

- `x = 10`: The **type is inferred** as `int` based on the value `10`.
- `name = "kira"`: The **type is inferred** as `string`.
- `z string = "random string"`: The **type is explicitly declared** as `string`.
- `y`: A **raw string literal**, which allows multi-line text and ignores escape sequences. It supports double quotes (`"`) without the need for backslashes.

---

## ✅ **2. Printing Values and Types:**
```go
fmt.Println(x)
fmt.Printf("Type: %T\n", x)
```
- `fmt.Println()` prints the value.
- `fmt.Printf()` with `%T` prints the **type of the variable**.

---

## ✅ **3. Static Typing in Go:**
```go
// name = 100
```
- Go is a **statically typed language**, meaning **the type of a variable is fixed at declaration**.
- If you try to assign an integer (`100`) to the variable `name` (which was inferred as a `string`), **the compiler will throw an error**.

---

## ✅ **4. Output:**
```
10
Type: int
kira
Type:  string
random string
Type: string
some string, "random string"
Type: string
```

---

## 🎯 **Key Concepts:**
| Concept             | Explanation                            |
|-----------------|----------------------------------|
| **Type Inference** | Go automatically determines the type based on the assigned value. |
| **Explicit Typing** | You can explicitly declare the type using `var z string = "text"`. |
| **Raw String Literal** | Uses backticks (` `` `) to include special characters and newlines. |
| **Static Typing** | Once a variable is assigned a type, it cannot be reassigned to a different type. |