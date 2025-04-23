This Go program generates the Fibonacci sequence up to a given number `n`. Here's a detailed breakdown of the code and how it works:

### **1. Understanding the Fibonacci Sequence**
The Fibonacci sequence is a series of numbers where each number (after the first two) is the sum of the two preceding ones. It typically starts with `0` and `1`. For example, the Fibonacci sequence up to `n = 10` would look like this:

```
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
```

### **2. Code Breakdown**

```go
package main

import "fmt"

// 2) ряд фибоначчи
// вводится какое либо число и небходимо посчитать число фибоначи до этого числа
// n=10
// 0 1 1 2 3 5 8 13
func main() {
	var n int
	fmt.Scanf("%d", &n)  // Read an integer input from the user and store it in the variable 'n'.
```
- **Explanation**: 
  - The program starts by importing the `fmt` package, which is used to print output and take user input.
  - The user is prompted to enter a number (`n`) up to which they want to calculate the Fibonacci sequence. `fmt.Scanf("%d", &n)` reads an integer value from the user and stores it in the variable `n`.

```go
	start1 := 0
	start2 := 1
```
- **Explanation**:
  - `start1` and `start2` represent the first two numbers in the Fibonacci sequence: `0` and `1`.
  
```go
	for true {
		c := start1 + start2  // Calculate the next Fibonacci number (sum of the previous two numbers)
		if c > n {  // If the next number exceeds 'n', stop the loop
			break
		}
		fmt.Println(c)  // Print the current Fibonacci number
		tmp := start1  // Store the value of 'start1' temporarily
		start1 = start2  // Update 'start1' to the previous 'start2'
		start2 = start1 + tmp  // Calculate the new 'start2' as the sum of 'start1' and the old 'start1'
	}
```
- **Explanation**:
  - The program enters an infinite loop (`for true`), and it keeps calculating Fibonacci numbers until it reaches a number greater than `n`.
  - **`c := start1 + start2`**: The next Fibonacci number is calculated by adding `start1` and `start2`.
  - **`if c > n`**: If the calculated Fibonacci number `c` exceeds `n`, the loop breaks, and the program stops calculating further Fibonacci numbers.
  - **`fmt.Println(c)`**: Prints the Fibonacci number `c` if it is less than or equal to `n`.
  - **`tmp := start1`**: Temporarily stores the value of `start1` to swap the values in the next iteration.
  - **`start1 = start2`**: Updates `start1` to the current value of `start2`.
  - **`start2 = start1 + tmp`**: Updates `start2` to the sum of the new `start1` and the previous value of `start1` (`tmp`), completing the update for the next Fibonacci number.

```go
}
```
- **Explanation**:
  - This marks the end of the `main` function.

### **3. Example Walkthrough**
Let’s go through an example with `n = 89`:

- User enters `n = 89`.
- Initially, `start1 = 0` and `start2 = 1`.
- In each loop iteration, the next Fibonacci number is calculated by adding `start1` and `start2`.

#### Loop 1:
- `c = start1 + start2 = 0 + 1 = 1`.
- Since `1 <= 89`, the program prints `1`.
- Update: `start1 = 1`, `start2 = 1` (because `start1 + start2 = 1`).

#### Loop 2:
- `c = 1 + 1 = 2`.
- Since `2 <= 89`, the program prints `2`.
- Update: `start1 = 1`, `start2 = 2`.

#### Loop 3:
- `c = 1 + 2 = 3`.
- Since `3 <= 89`, the program prints `3`.
- Update: `start1 = 2`, `start2 = 3`.

#### Loop 4:
- `c = 2 + 3 = 5`.
- Since `5 <= 89`, the program prints `5`.
- Update: `start1 = 3`, `start2 = 5`.

#### Loop 5:
- `c = 3 + 5 = 8`.
- Since `8 <= 89`, the program prints `8`.
- Update: `start1 = 5`, `start2 = 8`.

#### Loop 6:
- `c = 5 + 8 = 13`.
- Since `13 <= 89`, the program prints `13`.
- Update: `start1 = 8`, `start2 = 13`.

#### Loop 7:
- `c = 8 + 13 = 21`.
- Since `21 <= 89`, the program prints `21`.
- Update: `start1 = 13`, `start2 = 21`.

#### Loop 8:
- `c = 13 + 21 = 34`.
- Since `34 <= 89`, the program prints `34`.
- Update: `start1 = 21`, `start2 = 34`.

#### Loop 9:
- `c = 21 + 34 = 55`.
- Since `55 <= 89`, the program prints `55`.
- Update: `start1 = 34`, `start2 = 55`.

#### Loop 10:
- `c = 34 + 55 = 89`.
- Since `89 <= 89`, the program prints `89`.
- Update: `start1 = 55`, `start2 = 89`.

#### Loop 11:
- `c = 55 + 89 = 144`.
- Since `144 > 89`, the program breaks out of the loop.

### **4. Output**
For input `n = 89`, the output will be:
```
1
2
3
5
8
13
21
34
55
89
```

### **Summary**
- The program calculates and prints Fibonacci numbers starting from `1` and continues until it generates a number greater than `n`.
- The loop terminates when the next Fibonacci number exceeds the given `n`.
