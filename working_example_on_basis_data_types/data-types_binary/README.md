This Go program demonstrates how to format and print values using the `fmt.Printf` function with specific formatting verbs.

### Code Explanation:

```go
package main

import "fmt"

func main() {
	fmt.Printf("%d - %b \n", 21, 21)
}
```

### 1. **`fmt.Printf` Function:**
   The `fmt.Printf` function is used for formatted output in Go. It allows you to print values with specific formatting, which is controlled by formatting verbs (such as `%d`, `%b`, etc.).

### 2. **Formatting Verbs:**
   The two formatting verbs used in this example are:

   - **`%d`**: This verb formats the value as a **decimal integer** (base 10). So, when we pass `21` as an argument for `%d`, it will print the number in base 10.
   - **`%b`**: This verb formats the value as a **binary number** (base 2). When we pass `21` as an argument for `%b`, it will print the binary representation of `21`.

### 3. **The Print Statement:**
   ```go
   fmt.Printf("%d - %b \n", 21, 21)
   ```

   - The first argument (`21`) is formatted using `%d` (decimal), so it will print `21` in decimal (which is `21`).
   - The second argument (`21`) is formatted using `%b` (binary), so it will print the binary representation of `21`, which is `10101` (since `21` in binary is `10101`).

### 4. **Output:**
   The `fmt.Printf` will print:
   ```bash
   21 - 10101
   ```

### Breakdown of the Output:
   - **`%d`**: Formats the number `21` as a decimal integer (base 10).
   - **`%b`**: Formats the number `21` as a binary number (base 2). `21` in binary is `10101`.

### Summary:
- **`%d`**: Prints the value as a decimal integer.
- **`%b`**: Prints the value as a binary number.
- The `fmt.Printf` function formats and prints both the decimal and binary representations of the number `21`.