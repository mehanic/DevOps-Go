Let's break down the code you provided and explain how the loops work:

### **1. Simple `for` loop with a counter**

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

- This is a standard **for loop**.
- The loop initializes `i` to 0, and then runs while `i` is less than 10.
- On each iteration, the value of `i` is printed using `fmt.Println(i)`, and then `i` is incremented (`i++`).
- The loop stops when `i` reaches 10, and the numbers 0 through 9 will be printed to the console.

**Output:**
```
0
1
2
3
4
5
6
7
8
9
```

### **2. `for` loop with a condition (loop until sum reaches 1000)**

```go
sum := 0
for sum < 1000 {
    sum += 100
}
fmt.Println(sum)
```

- Here, we initialize a variable `sum` to 0.
- The loop continues running as long as `sum < 1000` (i.e., while `sum` is less than 1000).
- In each iteration, 100 is added to `sum` (`sum += 100`).
- The loop will run until `sum` becomes 1000 or greater.
- After the loop, `fmt.Println(sum)` prints the final value of `sum`, which will be `1000`.

**Output:**
```
1000
```

### **3. `for` loop with another counter (increments by 10)**

```go
k := 0
for k < 100 {
    k += 10
    fmt.Println(k)
}
```

- In this case, the variable `k` is initialized to 0.
- The loop runs while `k < 100`.
- In each iteration, 10 is added to `k` (`k += 10`), and the value of `k` is printed after each addition.
- This will print the values `10`, `20`, `30`, ..., up to `90`, and the loop will terminate when `k` is no longer less than 100.

**Output:**
```
10
20
30
40
50
60
70
80
90
```

### **4. Infinite loop examples**

#### **Infinite loop with a constant `true` condition**

```go
// without ending
// for true {
//    fmt.Println(11)
// }
```

- This is an **infinite loop** because the condition `true` is always true.
- The loop would print `11` indefinitely, creating an infinite loop that would never stop unless manually interrupted or a `break` statement is added inside the loop.

#### **Infinite loop with a `for` without condition**

```go
// for {
//    fmt.Println(11)
// }
```

- This is another form of an **infinite loop** in Go.
- A `for` loop without any condition is equivalent to `for true`.
- It will continue running indefinitely, printing `11` on each iteration.
- Again, it would run forever unless you break out of it or the program is manually stopped.

### **Summary of the Rules:**

1. **Standard `for` loop**: Used to repeat a block of code a specific number of times, with initialization, condition checking, and iteration.
2. **Loop with a condition**: The loop can continue until a certain condition is met (e.g., when `sum` becomes greater than or equal to 1000).
3. **Increments**: You can modify variables inside the loop (e.g., `sum += 100` or `k += 10`), which controls how the loop progresses.
4. **Infinite `for` loop**: A loop that will run indefinitely if the condition always evaluates to true. In Go, this is typically written as `for {}` or `for true {}`.
5. **Loop control**: In Go, a `for` loop can run indefinitely without an explicit condition, allowing for infinite loops.

In summary, the program demonstrates how to use basic loops in Go, including a standard loop, a loop that increments a variable until it meets a condition, and examples of infinite loops.