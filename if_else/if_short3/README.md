This Go program demonstrates the usage of `if`, `if-else`, `if-else if-else`, and `if` with an **initialization statement**. Let's break down each section.

---

## **1. Basic `if` Statement**
```go
if x > 5 {
    fmt.Println("x is greater than 5")
}
```
- The condition `x > 5` is evaluated.
- Since `x = 10`, the condition is `true`, so `"x is greater than 5"` is printed.
- If the condition was `false`, nothing would be printed.

---

## **2. `if-else` Statement**
```go
if x > 5 {
    fmt.Println("x is greater than 5")
} else {
    fmt.Println("x is not greater than 5")
}
```
- If `x > 5` is `true`, the first block executes.
- Otherwise, the `else` block executes.

Since `x = 10` and `10 > 5`, it prints:
```
x is greater than 5
```
If `x` were `3`, the output would be:
```
x is not greater than 5
```

---

## **3. `if-else if-else` Statement**
```go
if x > 10 {
    fmt.Println("x is greater than 10")
} else if x == 10 {
    fmt.Println("x is exactly 10")
} else {
    fmt.Println("x is less than 10")
}
```
- The program checks `if x > 10`. Since `x = 10`, this is `false`.
- It then checks `else if x == 10`, which is `true`, so `"x is exactly 10"` is printed.
- If neither condition was true, the `else` block would execute.

**Expected Output (since `x = 10`):**
```
x is exactly 10
```

---

## **4. `if` with Initialization Statement**
```go
if y := 20; y > 10 {
    fmt.Println("y is greater than 10")
}
```
- This declares and initializes `y` inside the `if` statement (`y := 20`).
- Then, it checks `if y > 10`, which is `true`, so `"y is greater than 10"` is printed.
- **Important:** `y` is only accessible inside the `if` block. It **does not exist outside**.

If you try:
```go
fmt.Println(y) // ERROR! y is not defined outside the if block.
```
You will get a **compile-time error** because `y` is out of scope.

---

## **Summary of Rules**
1. **`if` Statement** - Executes the block if the condition is `true`.
2. **`if-else` Statement** - If the `if` condition is `false`, the `else` block executes.
3. **`if-else if-else` Statement** - Checks multiple conditions in sequence. The first `true` condition executes, and the rest are ignored.
4. **`if` with Initialization** - You can declare a variable inside `if`. That variable is only valid within the `if` block.

### **Expected Output of the Program**
```
Basic if:
x is greater than 5
if..else:
x is greater than 5
if-else if-else:
x is exactly 10
With Initialization Statement:
y is greater than 10
```

Let me know if you need more clarification! 🚀