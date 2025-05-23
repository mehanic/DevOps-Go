### Explanation of the Go Program (Calculator App):

This Go program implements a basic calculator that can perform three operations: addition, division, and multiplication. It uses `bufio.Scanner` to read input from the user and `strconv.ParseInt` to convert the input into integers. Below is a detailed breakdown of the program and its functions.

### Code Breakdown:

#### 1. **Import Statements**:
```go
import "fmt"
import "os"
import "bufio"
import "strconv"
```
- `fmt`: Provides formatted I/O functions (e.g., `fmt.Println` and `fmt.Scan`).
- `os`: Used for interacting with the operating system, specifically to read input from `stdin`.
- `bufio`: Provides buffered I/O operations, allowing the program to efficiently read user input.
- `strconv`: Provides functions for converting strings to numeric types (e.g., `ParseInt` and `ParseFloat`).

#### 2. **Error Handling Functions**:
```go
func error_handler_int(err error) int64 {
    if err != nil {
        fmt.Println(err)
        return int64(-1)
    }
    return int64(0)
}

func error_handler_float(err error) float64 {
    if err != nil {
        fmt.Println(err)
        return float64(-1)
    }
    return float64(0)
}
```
- These functions handle errors that may occur when parsing strings into numeric types.
- If an error occurs (e.g., invalid input), an error message is printed, and a default "error value" (`-1`) is returned.
- If no error occurs, a value of `0` is returned.

#### 3. **Arithmetic Functions**:
The program implements three arithmetic functions for addition, division, and multiplication.

##### **Addition Function**:
```go
func add_two_number() int64 {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter the first number")
    scanner.Scan()
    a, err := strconv.ParseInt(scanner.Text(), 10, 64)
    error_handler_int(err)
    fmt.Println("Enter the second number")
    scanner.Scan()
    b, err := strconv.ParseInt(scanner.Text(), 10, 64)
    error_handler_int(err)
    return a + b
}
```
- This function reads two numbers from the user, converts them to integers using `strconv.ParseInt`, and then returns their sum.

##### **Division Function**:
```go
func divide_two_number() float64 {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter the first number")
    scanner.Scan()
    a, err := strconv.ParseInt(scanner.Text(), 10, 64)
    error_handler_float(err)
    fmt.Println("Enter the second number")
    scanner.Scan()
    b, err := strconv.ParseInt(scanner.Text(), 10, 64)
    error_handler_float(err)
    return float64(a) / float64(b)
}
```
- This function reads two numbers from the user, converts them to integers, and then returns their division as a floating-point number.
- The result is converted to a `float64` to handle fractional values (e.g., `5 / 2 = 2.5`).

##### **Multiplication Function**:
```go
func multiply_two_number() int64 {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter the first number")
    scanner.Scan()
    a, err := strconv.ParseInt(scanner.Text(), 10, 64)
    error_handler_int(err)
    fmt.Println("Enter the second number:")
    scanner.Scan()
    b, err := strconv.ParseInt(scanner.Text(), 10, 64)
    error_handler_int(err)
    return a * b
}
```
- This function reads two numbers, converts them to integers, and then returns their product.

#### 4. **Main Function**:
The `main` function serves as the entry point of the program and controls the flow of the calculator app.

```go
func main() {
    var num int64
    var num1 float64
    num = 0
    num1 = 0
    fmt.Println("Calculator app")
    fmt.Println("1:Add operation")
    fmt.Println("2:Substract operation")
    fmt.Println("3:Multiply operation")
    scanner1 := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter your choice:")
    scanner1.Scan()
    op, _ := strconv.ParseInt(scanner1.Text(), 10, 64)
    switch op {
        case 1:
            num = add_two_number()
            fmt.Println(num)
        case 2:
            num1 = divide_two_number()
            fmt.Println(num1)
        case 3:
            num = multiply_two_number()
            fmt.Println(num)
    }
}
```

- **Variables**:
  - `num`: Stores the result of integer operations (addition and multiplication).
  - `num1`: Stores the result of division (which is a floating-point number).

- **Displaying the Menu**:
  The program displays a simple menu for the user to choose an operation:
  ```
  1:Add operation
  2:Subtract operation
  3:Multiply operation
  ```

- **User Input for Operation**:
  The program reads the user's choice using `scanner1.Scan()`, which captures the user's input as a string. This input is then converted to an integer using `strconv.ParseInt`.

- **Switch Case for Operations**:
  Depending on the user's input (`op`), the program executes the corresponding arithmetic function:
  - `case 1`: Calls `add_two_number()` to add two numbers.
  - `case 2`: Calls `divide_two_number()` to divide two numbers.
  - `case 3`: Calls `multiply_two_number()` to multiply two numbers.

- **Output**: After each operation, the program prints the result of the operation (`num` or `num1`).

### Example Flow:

1. The program starts with:
   ```
   Calculator app
   1:Add operation
   2:Substract operation
   3:Multiply operation
   Enter your choice:
   ```
   - The user inputs `1` to choose addition.

2. The program then asks for the first and second numbers, one at a time:
   ```
   Enter the first number
   5
   Enter the second number
   3
   ```
   - The program will add `5` and `3`, and print the result:
   ```
   8
   ```

### Error Handling:
- When parsing the user input, if an error occurs (e.g., if the user enters non-numeric input), the `error_handler_int()` or `error_handler_float()` functions will print the error and return a default error value (`-1`), preventing the program from crashing.

### Conclusion:
This program implements a simple command-line calculator that can add, subtract, multiply, or divide two numbers. It handles user input errors gracefully and prints the result of the selected operation. The functions for each operation (addition, division, multiplication) are separated into their own functions for clarity and reusability.