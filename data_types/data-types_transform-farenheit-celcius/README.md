This Go program converts temperatures between **Celsius (°C)** and **Fahrenheit (°F)**. It consists of two parts:

---

## ✅ **1. `main` package (Command-line Argument Handling)**

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)
```

### 📌 The purpose of this section:
- Accepts **command-line arguments** (temperature values).
- Converts those arguments from **string to float64**.
- Converts the values between Celsius and Fahrenheit.
- Prints the conversion results.

---

### 🎯 **Step 1: Loop through the command-line arguments:**
```go
for _, arg := range os.Args[1:] {
```
- `os.Args[1:]` skips the program name and captures the arguments (temperatures).
- The loop iterates over each argument.

---

### 🎯 **Step 2: Convert the argument to `float64`:**
```go
t, err := strconv.ParseFloat(arg, 64)
if err != nil {
	fmt.Fprintf(os.Stderr, "cf: %v\n", err)
	os.Exit(1)
}
```
- `strconv.ParseFloat()` converts the argument to a `float64`.
- If the conversion fails (e.g., invalid input), an error message is printed, and the program exits.

---

### 🎯 **Step 3: Perform the conversion:**
```go
f := Fahrenheit(t)
c := Celsius(t)
```
- Convert the `float64` value `t` to **Celsius** and **Fahrenheit** types.

---

### 🎯 **Step 4: Print the conversion result:**
```go
fmt.Printf("%f = %f, %f = %f\n", f, FToC(f), c, CToF(c))
```
- Converts Fahrenheit to Celsius using `FToC()`.
- Converts Celsius to Fahrenheit using `CToF()`.
- Example output:
```
32.000000 = 0.000000, 0.000000 = 32.000000
```

---

## ✅ **2. `temperature` package (Conversion Logic and String Methods)**

```go
package main

import "fmt"
```
- Defines two **custom types**, `Celsius` and `Fahrenheit`.

---

### 🎯 **Constants for reference temperatures:**
```go
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)
```
- These are reference points for temperature conversion.

---

### 🎯 **Conversion functions:**
```go
func CToF(c Celsius) Fahrenheit { 
	return Fahrenheit(c*9/5 + 32) 
}

func FToC(f Fahrenheit) Celsius { 
	return Celsius((f - 32) * 5 / 9) 
}
```
- `CToF()`: Converts Celsius to Fahrenheit.
- `FToC()`: Converts Fahrenheit to Celsius.

---

### 🎯 **Custom `String()` methods for formatting output:**
```go
func (c Celsius) String() string { 
	return fmt.Sprintf("%g°C", c) 
}

func (f Fahrenheit) String() string { 
	return fmt.Sprintf("%g°F", f) 
}
```
- Overrides the default `String()` method.
- Allows printing the temperature with the appropriate symbol (°C or °F).

---

## ✅ **Output Example:**
```
$ go run main.go 32 100
32.000000 = 0.000000, 0.000000 = 32.000000
100.000000 = 37.777778, 100.000000 = 212.000000
```

---

## 🎯 **Key Concepts:**
| Concept               | Purpose |
|----------------|---------------------------|
| **Custom Types** | `Celsius` and `Fahrenheit` |
| **Conversion Functions** | `CToF()` and `FToC()` |
| **Command-line Parsing** | `os.Args` |
| **Error Handling** | Handle invalid input |
| **String Formatting** | `String()` methods for better output |

---

## ✅ **Final Summary:**
- Converts temperature between Celsius and Fahrenheit.
- Uses custom types and conversion functions for clarity.
- Implements error handling for invalid input.
- Provides clean and readable output with custom `String()` methods.