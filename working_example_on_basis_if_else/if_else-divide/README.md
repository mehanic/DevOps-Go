This Go program demonstrates a series of conditional checks that determine whether the number `a` (which is 16 in this case) is divisible by 8, 6, 4, or 2 without a remainder. Let's break it down:

### 1. **Variable Initialization**:
```go
a := 16
```
- This declares a variable `a` and initializes it with the value `16`.

### 2. **First `if` Condition (`a%8 == 0`)**:
```go
if a%8 == 0 {
    println("остатка при деление на 8 нету")
}
```
- **`a % 8`** calculates the remainder when `a` is divided by 8.
- **`a % 8 == 0`** checks if the remainder is `0`. If `a` is divisible by 8 without any remainder, this condition will be `true`.
- For `a = 16`, `16 % 8 == 0`, so the condition is `true`, and the message `"остатка при деление на 8 нету"` (which means "no remainder when divided by 8" in Russian) will be printed.
- Since the condition is met, the program **does not check the other conditions**.

### 3. **Other `else if` Conditions**:
```go
else if a%6 == 0 {
    println("остатка при деление на 6 нету")
} else if a%4 == 0 {
    println("остатка при деление на 4 нету")
} else if a%2 == 0 {
    println("остатка при деление на 2 нету")
}
```
- These conditions are only checked if the previous conditions (`a % 8 == 0`) were not satisfied.
- In this case, since the first condition (`a % 8 == 0`) was already met (for `a = 16`), none of these `else if` conditions are evaluated.

### 4. **The `else` Condition**:
```go
else {
    println("остаток есть всегда")
}
```
- The `else` block is executed if none of the previous conditions are satisfied.
- However, in this case, since the condition `a % 8 == 0` was true for `a = 16`, this `else` block is **not executed**.

### 5. **Final Output**:
Since `a = 16`, the program checks if it is divisible by 8. The condition `a % 8 == 0` is satisfied (since `16 % 8 == 0`), so the message `"остатка при деление на 8 нету"` is printed, which translates to "no remainder when divided by 8".

### **Output**:
```
остатка при деление на 8 нету
```

### Summary of the Rules:
1. **Checking Divisibility**: The program uses the modulus operator (`%`) to check if a number is divisible by another number without a remainder. 
   - If `a % x == 0`, then `a` is divisible by `x`.
   
2. **Order of Conditions**: The conditions are checked in the order they are written:
   - If the first condition (`a % 8 == 0`) is true, it immediately prints the corresponding message and skips the rest of the conditions.
   - If none of the previous conditions are true, it will fall back to the `else` block.

3. **Output**: Since `a = 16`, it meets the first condition (`a % 8 == 0`), so the output is `"остатка при деление на 8 нету"`.

### Explanation of the Russian Text:
- **"остатка при деление на 8 нету"**: "No remainder when divided by 8".
- **Other possible messages** (not printed in this case):
  - `"остатка при деление на 6 нету"`: "No remainder when divided by 6".
  - `"остатка при деление на 4 нету"`: "No remainder when divided by 4".
  - `"остатка при деление на 2 нету"`: "No remainder when divided by 2".
  - `"остаток есть всегда"`: "There is always a remainder". This would be printed if `a` is not divisible by 8, 6, 4, or 2.

In conclusion, this program checks for divisibility in descending order (from 8, 6, 4, 2) and prints a message accordingly. If none of the conditions are met, the program prints a default message.