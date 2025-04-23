Let's break down the provided Go code to understand how it works.

### Code Overview:

```go
package main

import "fmt"

func sum(n int) int {
	var x int = 0
	for n > 0 {
		x += n % 10
		n = n / 10
	}
	return x
}

func numbers_sum(n int) int {
	var x int = 0
	for n > 0 {
		x += n % 10
		n = n / 10
	}
	if x < 10 {
		return x
	}
	return numbers_sum(x)  
}

func main() {
	fmt.Println(numbers_sum(456))
}
```

### Explanation:

1. **`sum` function**:
    ```go
    func sum(n int) int {
        var x int = 0
        for n > 0 {
            x += n % 10  // Adds the last digit of `n` to `x`
            n = n / 10   // Removes the last digit from `n`
        }
        return x
    }
    ```
    - The `sum` function calculates the sum of the digits of a number `n`.
    - **Logic**: The function uses a loop to extract each digit of the number `n` by using the modulo operation (`n % 10`). It adds this digit to the variable `x`. Then, the number `n` is reduced by removing the last digit using integer division (`n / 10`). This process continues until `n` becomes 0.
    - For example:
      - For `n = 456`, the sum of the digits is `4 + 5 + 6 = 15`.

2. **`numbers_sum` function**:
    ```go
    func numbers_sum(n int) int {
        var x int = 0
        for n > 0 {
            x += n % 10  // Adds the last digit of `n` to `x`
            n = n / 10   // Removes the last digit from `n`
        }
        if x < 10 {
            return x  // If the sum of digits is less than 10, return the result.
        }
        return numbers_sum(x)  // Otherwise, recursively calculate the sum of digits.
    }
    ```
    - This function aims to calculate the "digital root" of a number. The digital root is the final sum of digits of a number, which is obtained by repeatedly summing the digits until a single digit is reached.
    - **Logic**:
      - The function sums the digits of the number `n`, just like the `sum` function.
      - After calculating the sum (`x`), it checks if `x` is less than 10. If yes, it returns `x` (since the result is already a single-digit number).
      - If `x` is 10 or greater, the function calls itself recursively with `x` as the new input. This recursion keeps happening until a single-digit result is obtained.
      
      - For example:
        - For `n = 456`, the sum of digits is `4 + 5 + 6 = 15`. Since `15` is greater than 10, the function calls itself with `15`.
        - For `n = 15`, the sum of digits is `1 + 5 = 6`. Since `6` is less than 10, the function returns `6`.

3. **`main` function**:
    ```go
    func main() {
        fmt.Println(numbers_sum(456))
    }
    ```
    - The `main` function calls the `numbers_sum` function with the number `456` and prints the result.
    - As explained above, for `456`, the sum of digits is `15`, and then the sum of digits of `15` is `6`. Hence, the program will print `6`.

### Key Concepts:

- **Digital Root**: This is the result of repeatedly summing the digits of a number until a single-digit number is obtained.
  - The sum of the digits of `456` is `15`, and the sum of the digits of `15` is `6`, so the digital root of `456` is `6`.

### Example Walkthrough:

1. **First Call (`numbers_sum(456)`)**:
    - Sum the digits: `4 + 5 + 6 = 15`.
    - Since `15` is greater than `10`, we call `numbers_sum(15)`.

2. **Second Call (`numbers_sum(15)`)**:
    - Sum the digits: `1 + 5 = 6`.
    - Since `6` is less than `10`, we return `6`.

3. The final result is `6`, which is printed by the `main` function.

### Simplified Explanation of the Code Flow:

- The `numbers_sum` function sums the digits of a number repeatedly until a single-digit result is achieved.
- The recursion stops when the sum of digits is less than 10.
- In this example, the sum of digits of `456` is `15`, and then the sum of digits of `15` is `6`, which is the final result.

### Performance Consideration:
- For large numbers, the function works efficiently by breaking down the number into smaller components (i.e., its digits) and performing the summation operation, which reduces the number to a single digit in a few steps.