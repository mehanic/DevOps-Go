### Explanation of the Go Program

This Go program iterates from `1` to `100`, and prints only the **odd numbers** by skipping the even numbers. Let's break down the code:

---

### 1. **The `for` Loop:**

```go
for n := 1; n <= 100; n++ {
```

- This `for` loop initializes the variable `n` to `1` and runs the loop until `n` is less than or equal to `100` (`n <= 100`).
- On each iteration of the loop, the value of `n` is incremented by 1 (`n++`), so `n` will take values from `1` to `100`.

---

### 2. **Skipping Even Numbers with `continue`:**

```go
if n%2 == 0 {
    continue
}
```

- This `if` statement checks if `n` is an **even number**. The condition `n%2 == 0` means that if the remainder when `n` is divided by `2` is `0`, then `n` is an even number.
  
  For example:
  - `n = 2` → `2 % 2 == 0` → Even number
  - `n = 3` → `3 % 2 == 1` → Odd number

- If the condition is true (i.e., `n` is even), the **`continue`** statement is executed. The `continue` statement tells the program to skip the rest of the code inside the loop for that particular iteration and move to the next iteration.
  
  This means the `fmt.Println(n)` statement will **not** be executed for even numbers.

---

### 3. **Printing Odd Numbers:**

```go
fmt.Println(n)
```

- If the number `n` is **odd** (i.e., when `n%2 != 0`), the `fmt.Println(n)` statement will execute, and it will print the current value of `n`.

- Since the even numbers are skipped due to the `continue` statement, only the odd numbers from `1` to `100` will be printed.

---

### 4. **Output:**

The output will print only the odd numbers from `1` to `100`:

```
1
3
5
7
9
11
13
15
17
19
21
23
25
27
29
31
33
35
37
39
41
43
45
47
49
51
53
55
57
59
61
63
65
67
69
71
73
75
77
79
81
83
85
87
89
91
93
95
97
99
```

---

### **Key Points:**

1. **Loop Initialization**: The loop starts with `n = 1` and runs until `n <= 100`.
2. **Even Numbers Check**: The `if n%2 == 0` condition checks whether the current number `n` is even.
3. **Skipping Even Numbers**: If `n` is even, the `continue` statement skips the rest of the loop's body (including `fmt.Println`).
4. **Printing Odd Numbers**: Only when `n` is odd (`n % 2 != 0`), the number gets printed.

---

This program effectively demonstrates how to **skip even numbers** and **only print odd numbers** using the `continue` statement. Would you like to modify this logic or add any other functionality?