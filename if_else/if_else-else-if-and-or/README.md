This Go program demonstrates a series of conditional structures (`if`, `else`, `else if`, and the use of logical operators like `&&` (AND) and `||` (OR)) to check various conditions and print appropriate messages based on those conditions. Let's break it down:

### **1. `main` function (Using Simple `if` Statements)**:
```go
func main() {
    number := 50
    var grade string

    if number <= 50 {
        grade = "Bad" // If number is less than or equal to 50, grade is set to "Bad"
    }

    if number >= 51 {
        grade = "Good" // If number is greater than or equal to 51, grade is set to "Good"
    }

    fmt.Println(grade) // Prints the grade

    main1()
    main3()
    main4()
}
```
- **Logic**:
  - `if number <= 50` checks if the `number` is less than or equal to 50. If true, it assigns `"Bad"` to `grade`.
  - The next `if` statement (`if number >= 51`) checks if `number` is greater than or equal to 51. If true, it assigns `"Good"` to `grade`.
  - **Note**: Since both conditions are independent and `number` is 50, it first sets `grade` to `"Bad"` because of the first condition. However, it then reassigns `grade` to `"Good"` because of the second condition. Thus, the output will be `"Good"`.

- **Output of `main`**: 
  ```
  Good
  ```

### **2. `main1` function (Using `if` and `else`)**:
```go
func main1() {
    number := 70
    if number <= 50 {
        fmt.Println("Failed!") // If number <= 50, print "Failed!"
    } else {
        fmt.Println("Good Job!") // Otherwise, print "Good Job!"
    }
}
```
- **Logic**:
  - This is a simple `if-else` statement.
  - Since `number` is `70`, the condition `number <= 50` is false, so the `else` block is executed, and `"Good Job!"` is printed.

- **Output of `main1`**:
  ```
  Good Job!
  ```

### **3. `main2` function (Using `else if`)**:
```go
func main2() {
    number := 90
    if number <= 50 {
        fmt.Println("Failed!") // If number <= 50, print "Failed!"
    } else if number <= 70 {
        fmt.Println("Good Job!") // If number <= 70, print "Good Job!"
    } else {
        fmt.Println("Perfect!") // Otherwise, print "Perfect!"
    }
}
```
- **Logic**:
  - The first condition checks if `number <= 50`. Since `90` is not less than or equal to 50, it moves to the next condition.
  - The second condition checks if `number <= 70`. Since `90` is greater than 70, it moves to the `else` block.
  - The `else` block prints `"Perfect!"`.

- **Output of `main2`**:
  ```
  Perfect!
  ```

### **4. `main3` function (Using `&&` (AND) and `||` (OR) Operators)**:
```go
func main3() {
    number1 := 50
    if number1 <= 50 && number1 > 0 { // If number1 is between 1 and 50 inclusive
        fmt.Println("Not Passed") // Print "Not Passed"
    }

    number2 := 70
    if number2 <= 70 || number2 >= 51 { // If number2 is <= 70 or >= 51
        fmt.Println("Passed") // Print "Passed"
    }

    number3 := 99
    if number3 > 70 && number3 <= 100 { // If number3 is between 71 and 100 inclusive
        fmt.Println("Passed and won a prize") // Print "Passed and won a prize"
    }

    number4 := 10
    if number4 != 0 { // If number4 is not 0
        fmt.Println("Still have a chance") // Print "Still have a chance"
    }
}
```
- **Logic**:
  - `number1 <= 50 && number1 > 0`: The condition checks if `number1` is between 1 and 50. Since `50` is within this range, it prints `"Not Passed"`.
  - `number2 <= 70 || number2 >= 51`: The condition checks if `number2` is either less than or equal to 70 or greater than or equal to 51. Since `70` satisfies the condition, it prints `"Passed"`.
  - `number3 > 70 && number3 <= 100`: The condition checks if `number3` is between 71 and 100. Since `99` satisfies the condition, it prints `"Passed and won a prize"`.
  - `number4 != 0`: The condition checks if `number4` is not equal to zero. Since `10` is not zero, it prints `"Still have a chance"`.

- **Output of `main3`**:
  ```
  Not Passed
  Passed
  Passed and won a prize
  Still have a chance
  ```

### **5. `main4` function (Using `if` Short Statement and String Comparison)**:
```go
func main4() {
    firstName := "Prasetiyo"

    // Simple if-else based on a string comparison
    if firstName == "Prasetiyo" {
        fmt.Println("Hello,", firstName) // If firstName == "Prasetiyo", print "Hello, Prasetiyo"
    } else if firstName == "Yosho" {
        fmt.Println("Hello Yosho") // If firstName == "Yosho", print "Hello Yosho"
    } else {
        var input string
        fmt.Println("What is your name?") // If firstName is neither, prompt the user
        fmt.Scan(&input)
        fmt.Println("- Nice to meet you", input)
    }

    // Short If Statement
    if length := len(firstName); length < 5 {
        fmt.Println("Your name is too short") // If length < 5, print this
    } else if length < 10 {
        fmt.Println("Your name is just right") // If length < 10, print this
    } else {
        fmt.Println("Your name is too long") // Otherwise, print this
    }
}
```
- **Logic**:
  - **First Part (String Comparison)**: It checks if `firstName` is `"Prasetiyo"`. Since it is, it prints `"Hello, Prasetiyo"`. If `firstName` had been `"Yosho"`, it would print `"Hello Yosho"`. Otherwise, it would ask for the user's name.
  - **Second Part (Short If Statement)**: The `len(firstName)` is calculated inline using the `if` statement. Since `firstName` has 9 characters, it prints `"Your name is just right"` because the length is between 5 and 10.

- **Output of `main4`**:
  ```
  Hello, Prasetiyo
  Your name is just right
  ```

### **Final Output**:
```
Good
Good Job!
Perfect!
Not Passed
Passed
Passed and won a prize
Still have a chance
Hello, Prasetiyo
Your name is just right
```

### **Summary of Key Concepts**:
1. **`if`, `else if`, and `else`**: These structures allow you to execute different code based on certain conditions. You can have multiple `else if` conditions, and only the first true condition is executed.
   
2. **Logical Operators**:
   - `&&` (AND): Both conditions must be true for the entire expression to be true.
   - `||` (OR): Only one condition needs to be true for the entire expression to be true.
   - `!=` (Not equal to): Checks if two values are not equal.
   
3. **Short If Statements**: Go allows you to declare and initialize variables inline within the `if` statement, such as `if length := len(firstName); length < 5`.
   
4. **String Comparison**: You can compare strings directly using `==`, as shown in the first `if` condition.

This program demonstrates various ways to check conditions and execute different actions based on the results of those checks.