This Go program demonstrates how to print the binary representation of a number using `fmt.Printf()` with different formatting verbs. Let's break down what each part of the code does:

### Code Breakdown:

```go
fmt.Printf("%d : %b\n", 35, 35)
```

1. **`%d`**: This is the **decimal format** verb. It prints the integer in **base 10** (decimal form).
   - Here, `35` will be printed as `35` in decimal.

2. **`%b`**: This is the **binary format** verb. It prints the integer in **base 2** (binary form).
   - Here, `35` will be printed in its binary form, which is `100011`.

So, when `fmt.Printf("%d : %b\n", 35, 35)` is executed, the output will be:
```
35 : 100011
```

### Second `fmt.Printf` Call:

```go
fmt.Printf("%d : %#b\n", 35, 35)
```

1. **`%d`**: Again, this prints the integer in **base 10** (decimal). So, `35` will be printed as `35`.

2. **`%#b`**: The `#` flag in formatting verbs is used to add a prefix for certain types of output. When used with **binary format (`%b`)**, it will add the **`0b`** prefix before the binary representation.
   - For the number `35`, it will be printed as `0b100011`.

So, when `fmt.Printf("%d : %#b\n", 35, 35)` is executed, the output will be:
```
35 : 0b100011
```

### Summary of Rules:

1. **`%d`**: Prints the number in **decimal (base 10)**.
2. **`%b`**: Prints the number in **binary (base 2)**.
3. **`%#b`**: Prints the number in **binary** but adds a `0b` prefix to indicate it is in binary format.

### Final Output:
```
35 : 100011
35 : 0b100011
```