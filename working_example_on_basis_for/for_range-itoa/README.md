Let's break down the Go code step by step and explain each part in detail:

### Code Breakdown

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//strconv.Itoa-> numbers to string
	//strconv.Atoi-> string to numbers
	//рекурсию
	n := 2100
	primeNumbers := []string{}
	for i := 2; i < n; i++ {
		counter := 0
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				counter += 1
			}
		}
		if counter == 1 {
			primeNumbers = append(primeNumbers, strconv.Itoa(i))
		}
	}
	s := strings.Join(primeNumbers, "")
	var c int
	arr := []int{}
	fmt.Scan(&c)
	for i := 0; i < c; i++ {
		var d int
		fmt.Scan(&d)
		arr = append(arr, d)
	}
	for _, v := range arr {
		fmt.Print(string(s[v-1]))
	}
}
```

### Explanation of Each Part

#### 1. **Importing Packages**:
```go
import (
	"fmt"
	"strconv"
	"strings"
)
```
- `fmt`: This package is used for formatted I/O (e.g., printing output and scanning input).
- `strconv`: This package provides functions for converting between strings and numeric types. In this code, it's used for converting numbers to strings (`strconv.Itoa()`) and strings to numbers (`strconv.Atoi()`).
- `strings`: This package provides functions for manipulating strings. Specifically, it is used here to join a slice of strings into a single string (`strings.Join()`).

#### 2. **Generating Prime Numbers (up to 2100)**:
```go
n := 2100
primeNumbers := []string{}
for i := 2; i < n; i++ {
	counter := 0
	for j := 2; j <= i; j++ {
		if i%j == 0 {
			counter += 1
		}
	}
	if counter == 1 {
		primeNumbers = append(primeNumbers, strconv.Itoa(i))
	}
}
```
- `n := 2100`: We are interested in finding prime numbers up to `2100`.
- `primeNumbers := []string{}`: This initializes an empty slice (`primeNumbers`) that will store the prime numbers as strings.
- The outer `for` loop (`for i := 2; i < n; i++`) iterates over all numbers from `2` to `2099`.
  - For each number `i`, an inner loop checks if `i` is divisible by any number between `2` and `i` itself. This is essentially checking if `i` is prime.
  - `counter := 0`: This keeps track of how many divisors the number `i` has. If `i` is prime, it should only be divisible by `1` and itself (so `counter` should be `1`).
  - `if counter == 1`: If the number `i` has exactly one divisor (itself), then it is prime, and we convert it to a string (`strconv.Itoa(i)`) and append it to the `primeNumbers` slice.

#### 3. **Joining Prime Numbers into a Single String**:
```go
s := strings.Join(primeNumbers, "")
```
- `strings.Join(primeNumbers, "")`: This joins all the prime numbers in the `primeNumbers` slice into a single string `s`. The `""` means that there will be no separator between the prime numbers. This means that all the prime numbers will be concatenated into one long string.
- For example, if the prime numbers are `[2, 3, 5]`, the result of `strings.Join(primeNumbers, "")` would be the string `"235"`.

#### 4. **Reading User Input**:
```go
var c int
arr := []int{}
fmt.Scan(&c)
for i := 0; i < c; i++ {
	var d int
	fmt.Scan(&d)
	arr = append(arr, d)
}
```
- `var c int`: A variable `c` is declared to store how many numbers the user wants to input.
- `arr := []int{}`: This initializes an empty slice `arr` to store the user inputs.
- `fmt.Scan(&c)`: This reads an integer from the user and stores it in `c`, indicating how many values the user will input.
- The `for` loop runs `c` times and prompts the user to input an integer `d` each time (`fmt.Scan(&d)`), appending each value to the `arr` slice.

#### 5. **Printing Characters from the Prime Number String**:
```go
for _, v := range arr {
	fmt.Print(string(s[v-1]))
}
```
- This loop iterates over the slice `arr`, where each element `v` represents an index.
  - `s[v-1]`: This accesses the character at the index `v-1` in the string `s` (because Go slices and strings are zero-indexed, so we subtract `1` from `v` to get the correct index).
  - `fmt.Print(string(s[v-1]))`: This converts the character at index `v-1` to a string and prints it.

### Example Walkthrough:

#### Let's assume:
- `n = 2100` is the upper limit for finding prime numbers.
- The prime numbers generated between 2 and 2100 are stored in `primeNumbers`.
- The string `s` might look something like `"2357111317..."` (a concatenation of prime numbers).
- Suppose the user enters `3` for `c`, meaning they want to print 3 characters.
- The user then enters the values `1`, `5`, and `7` (for example).

The loop `for _, v := range arr` will:
1. For `v = 1`, print `s[1-1]` which corresponds to the first character of `s` (the first prime number, "2").
2. For `v = 5`, print `s[5-1]` which corresponds to the 4th character in `s`.
3. For `v = 7`, print `s[7-1]` which corresponds to the 6th character in `s`.

### Summary of the Program:

1. It generates a list of prime numbers from 2 up to 2100 and concatenates them into a string.
2. It then asks the user for how many numbers (`c`) they want to enter.
3. The user inputs `c` indices, and the program retrieves and prints the characters from the concatenated prime numbers string at those indices.

### Key Points:
- **Prime number generation**: The program generates prime numbers by checking divisibility.
- **String manipulation**: The prime numbers are concatenated into a single string and indexed.
- **User interaction**: The program asks for multiple indices and prints characters from the prime numbers string based on user input.