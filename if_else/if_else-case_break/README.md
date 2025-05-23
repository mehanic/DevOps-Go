This Go program is designed to randomly generate a number between 0 and 99, then determine whether that number is odd or even, and finally print a message based on the result.

### Code Explanation

Let's break the code down step by step:

### 1. **Imports**:
```go
import (
	"fmt"
	"math/rand"
	"time"
)
```
- **`fmt`**: This package is used for formatted I/O operations, like printing to the console.
- **`math/rand`**: This package is used to generate random numbers.
- **`time`**: This package is used to work with time, specifically for generating a source for randomness using the current Unix timestamp (`time.Now().UnixNano()`).

### 2. **Global Variable**:
```go
var st string
```
- A global variable `st` is declared of type `string`. This variable will hold the result of whether a number is odd or even, which is set in the `oddOrEven()` function and returned to the `main()` function.

### 3. **Main Function**:
```go
func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n1 := r1.Intn(100)
```
- **`s1 := rand.NewSource(time.Now().UnixNano())`**: This line creates a new random source based on the current Unix timestamp in nanoseconds. This ensures that the random number generator is different every time the program runs (because the timestamp changes).
  
- **`r1 := rand.New(s1)`**: This creates a new random number generator (`r1`) using the random source `s1`.

- **`n1 := r1.Intn(100)`**: This generates a random integer `n1` between `0` and `99` (because `Intn(n)` generates a number in the range `[0, n)`).

### 4. **Odd or Even Check**:
```go
stuation := oddOrEven(n1 % 2)
fmt.Println(n1, "is", stuation)
```
- **`n1 % 2`**: The `%` operator calculates the remainder when `n1` is divided by `2`. If `n1 % 2 == 0`, the number is even; if `n1 % 2 == 1`, the number is odd.

- **`stuation := oddOrEven(n1 % 2)`**: This calls the `oddOrEven()` function with the result of `n1 % 2`. It will pass either `0` (even) or `1` (odd) to the function.

- **`fmt.Println(n1, "is", stuation)`**: This prints the result to the console, showing the number (`n1`) and whether it is "an Odd Number" or "an Even Number".

### 5. **The `oddOrEven` Function**:
```go
func oddOrEven(num int) string {

	switch num {
	case 0:
		st = "an Even Number"
		break
	case 1:
		st = "an Odd Number"
		break
	}
	return st
}
```
- **`switch num`**: This is a `switch` statement that checks the value of `num` (which is either `0` or `1`).
  
- **`case 0:`**: If `num` is `0`, it means the number is even. So, the global variable `st` is set to `"an Even Number"`.
  
- **`case 1:`**: If `num` is `1`, it means the number is odd. So, the global variable `st` is set to `"an Odd Number"`.
  
- **`break`**: The `break` statement is unnecessary in this case since the `switch` statement will exit as soon as a `case` is matched. However, it is included here for clarity.

- **`return st`**: The function returns the value of `st`, which is either `"an Even Number"` or `"an Odd Number"`.

### 6. **Example Output**:
If the randomly generated number is `55`:
- `n1 = 55`
- `n1 % 2 = 1` (since `55 % 2` equals `1`), so the `oddOrEven()` function returns `"an Odd Number"`.
- The output would be:
  ```
  55 is an Odd Number
  ```

If the randomly generated number is `42`:
- `n1 = 42`
- `n1 % 2 = 0` (since `42 % 2` equals `0`), so the `oddOrEven()` function returns `"an Even Number"`.
- The output would be:
  ```
  42 is an Even Number
  ```

### Summary of Rules:
1. **Random Number Generation**: A random number `n1` between `0` and `99` is generated using `rand.NewSource` and `rand.New`.
2. **Odd or Even Check**: The number's odd or even status is determined using the remainder operation (`n1 % 2`).
3. **`oddOrEven()` Function**: This function uses a `switch` statement to check if the number is odd (`1`) or even (`0`) and returns the corresponding string message ("an Odd Number" or "an Even Number").
4. **Result Output**: The result (random number and its status) is printed to the console.

### Output Example:
```
77 is an Odd Number
```

Or:
```
24 is an Even Number
```

The output will vary every time you run the program due to the random number generation.