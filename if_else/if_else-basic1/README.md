This Go code contains various examples that demonstrate conditional logic (`if`, `else`, `else if`) and argument parsing with `os.Args`, as well as some constant usage. Let's break down the rules and explain what each block of code does:

### 1. **Main function (main)**

```go
func main() {
    score, valid := 5, true

    if score > 3 && valid {
        fmt.Println("good")
    }
}
```
- **Explanation**: This code checks two conditions:
  - If the `score` is greater than 3.
  - If the `valid` variable is `true`.
- If both conditions are true, it prints "good". In this case, `score` is 5, and `valid` is `true`, so it prints "good".

### 2. **Main function (main1)**

```go
func main1() {
    score, valid := 3, true

    if score > 3 && valid {
        fmt.Println("good")
    } else {
        fmt.Println("low")
    }
}
```
- **Explanation**: Here, the program checks if `score` is greater than 3 and `valid` is `true`. If the condition is true, it prints "good". Otherwise, it prints "low".
- Since `score` is 3, the condition `score > 3` is false, so it prints "low".

### 3. **Main function (main3)**

```go
func main3() {
    score := 3

    if score > 3 {
        fmt.Println("good")
    } else if score == 3 {
        fmt.Println("on the edge")
    } else {
        fmt.Println("low")
    }
}
```
- **Explanation**: This block uses `if`, `else if`, and `else` to handle different cases for the `score`:
  - If the score is greater than 3, it prints "good".
  - If the score is exactly 3, it prints "on the edge".
  - Otherwise, it prints "low".
- Since `score` is 3, it will print "on the edge".

### 4. **Main function (main4)**

```go
func main4() {
    score := 2

    if score > 3 {
        fmt.Println("good")
    } else if score == 3 {
        fmt.Println("on the edge")
    } else if score == 2 {
        fmt.Println("meh...")
    } else {
        fmt.Println("low")
    }
}
```
- **Explanation**: This block is similar to `main3` but has an additional condition for `score == 2`, where it prints "meh...".
- Since `score` is 2, it will print "meh...".

### 5. **Main function (main5)**: **Command-Line Argument Parsing & Conversion**

```go
const usage = `
Feet to Meters
--------------
This program converts feet into meters.

Usage:
feet [feetsToConvert]`

func main5() {
    if len(os.Args) < 2 {
        fmt.Println(strings.TrimSpace(usage))
        return
    }

    arg := os.Args[1]

    feet, _ := strconv.ParseFloat(arg, 64)

    meters := feet * 0.3048

    fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
```
- **Explanation**: This program converts feet to meters using command-line arguments. It checks if the user provided at least one argument (`len(os.Args) < 2`). If not, it prints a usage message.
- The program takes the first argument (`os.Args[1]`), converts it to a float using `strconv.ParseFloat`, and then multiplies it by `0.3048` to convert feet to meters.
- If the user provides a valid number (e.g., `5`), it will convert it to meters and print the result.

### 6. **Main function (main6)**: **Access Control Example**

```go
func main6() {
    args := os.Args

    if len(args) != 3 {
        fmt.Println("Usage: [username] [password]")
        return
    }

    u, p := args[1], args[2]

    if u != "jack" {
        fmt.Printf("Access denied for %q.\n", u)
    } else if p != "1888" {
        fmt.Printf("Invalid password for %q.\n", u)
    } else {
        fmt.Printf("Access granted to %q.\n", u)
    }
}
```
- **Explanation**: This program performs a simple access control check:
  - It expects two arguments: a username and a password.
  - If the username is not "jack", it denies access.
  - If the username is "jack" but the password is incorrect, it prints "Invalid password".
  - If both the username and password match, it grants access.
  
### 7. **Main function (main7)**: **Constant Values for Better Maintainability**

```go
const (
    usage1   = "Usage: [username] [password]"
    errUser  = "Access denied for %q.\n"
    errPwd   = "Invalid password for %q.\n"
    accessOK = "Access granted to %q.\n"
    user     = "jack"
    pass     = "1888"
)

func main7() {
    args := os.Args

    if len(args) != 3 {
        fmt.Println(usage1)
        return
    }

    u, p := args[1], args[2]

    if u != user {
        fmt.Printf(errUser, u)
    } else if p != pass {
        fmt.Printf(errPwd, u)
    } else {
        fmt.Printf(accessOK, u)
    }
}
```
- **Explanation**: This example is similar to `main6`, but it uses constants for values like `user`, `pass`, and messages, making the code more maintainable.
- If the username or password doesn't match the constants, the program will output an appropriate error message.

### 8. **Main function (main8)**: **Multiple Users and Passwords**

```go
const (
    usage0       = "Usage: [username] [password]"
    errUser0     = "Access denied for %q.\n"
    errPwd0      = "Invalid password for %q.\n"
    accessOK0    = "Access granted to %q.\n"
    user0, user2 = "jack", "inanc"
    pass0, pass2 = "1888", "1879"
)

func main8() {
    args := os.Args

    if len(args) != 3 {
        fmt.Println(usage)
        return
    }

    u, p := args[1], args[2]

    if u != user && u != user2 {
        fmt.Printf(errUser, u)
    } else if u == user && p == pass {
        fmt.Printf(accessOK, u)
    } else if u == user2 && p == pass2 {
        fmt.Printf(accessOK, u)
    } else {
        fmt.Printf(errPwd, u)
    }
}
```
- **Explanation**: This block extends the previous example to allow multiple valid users with different passwords.
  - If the username is not one of the valid users (`jack` or `inanc`), access is denied.
  - If the username matches but the password is incorrect, it prints an invalid password message.
  - If both the username and password match, it grants access.

### Summary of Rules:
1. **Conditionals**: Use `if`, `else if`, and `else` to handle different conditions and print appropriate messages.
2. **Command-Line Arguments**: Use `os.Args` to parse arguments and check for proper input length.
3. **Constants**: Define constants for values that are used multiple times, such as usernames and error messages, for better readability and maintainability.
4. **Access Control**: Use conditional checks to validate user credentials.