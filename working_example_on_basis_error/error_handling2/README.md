Your Go program demonstrates **custom error handling** by implementing the `error` interface for different types of errors. However, there are some **mistakes** that need to be fixed.

---

## **Key Rules in the Code**
### **1️⃣ Implementing Custom Errors**
Each custom error type must implement the `Error() string` method, which returns a descriptive error message.

#### **🔹 Incorrect Custom Error Implementations**
```go
func (e SquareNegativeSide) Error() (s string) {
    s = fmt.Sprintf("error, side can't be negative!! You input: %v", float64(e*e))
    return
}
```
```go
func (e ErrNegativeSqrt) Error() (s string) {
    s = fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
    return
}
```
✅ **Fix**: The `Error()` method should return a string **without named return values**:
```go
func (e SquareNegativeSide) Error() string {
    return fmt.Sprintf("error, side can't be negative!! You input: %v", float64(e))
}
```
```go
func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}
```
---
### **2️⃣ Using Custom Errors Correctly**
#### **🔹 Incorrect Error Assignment in `Sqrt()` and `Area()`**
```go
func Sqrt(x float64) (float64, error) {
    var n ErrNegativeSqrt
    var e error
    if x < 0 {
        e = n  // ❌ `n` is not initialized with a value!
        return x, e
    }
    return x, nil
}
```
✅ **Fix**: Assign a proper value to `n` before returning it.
```go
func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x) // ✅ Correct way to return custom error
    }
    return x, nil
}
```
---
### **3️⃣ Handling Errors in `Area()`**
#### **🔹 Incorrect Error Assignment**
```go
func Area(x float64) (float64, error) {
    var y SquareNegativeSide
    var e error
    if x < 0 {
        e = y // ❌ `y` is not initialized
        return x, e
    }
    return x * x, nil
}
```
✅ **Fix**: Assign a proper value to `y`:
```go
func Area(x float64) (float64, error) {
    if x < 0 {
        return 0, SquareNegativeSide(x) // ✅ Correct way to return custom error
    }
    return x * x, nil
}
```
---
### **4️⃣ Custom Errors with Struct (`AgeError`)**
#### **🔹 Issue: `Error()` Method Returns Named Value**
```go
func (er *AgeError) Error() (s string) { // ❌ Named return value is unnecessary
    s = fmt.Sprintf("age: %v\n message:%v", er.age, er.what)
    return s
}
```
✅ **Fix**: Return the string directly:
```go
func (er *AgeError) Error() string {
    return fmt.Sprintf("age: %v\nmessage: %v", er.age, er.what)
}
```
---
### **5️⃣ Handling `Person.getAge()` Correctly**
#### **🔹 Problem: Unnecessary Variable Assignments**
```go
func (per *Person) getAge() (int, error) {
    var myErr AgeError
    var e error
    myErr.age = per.age
    myErr.what = "age can't be <0 or >100"

    if per.age < 0 || per.age > 100 {
        e = &myErr
        myErr.age = per.age // ❌ Unnecessary reassignment
        myErr.what = "age can't be <0 or >100"
        return per.age, e
    }
    return per.age, nil
}
```
✅ **Fix**: Return the error directly:
```go
func (per *Person) getAge() (int, error) {
    if per.age < 0 || per.age > 100 {
        return per.age, &AgeError{per.age, "age can't be <0 or >100"} // ✅ Directly creating an instance
    }
    return per.age, nil
}
```
---
## **🔹 Final Corrected Code**
```go
package main

import (
	"fmt"
)

// Custom error for negative square root
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Custom error for negative square area
type SquareNegativeSide float64

func (e SquareNegativeSide) Error() string {
	return fmt.Sprintf("error, side can't be negative! You input: %v", float64(e))
}

// Custom error for invalid age
type AgeError struct {
	age  int
	what string
}

func (er *AgeError) Error() string {
	return fmt.Sprintf("age: %v\nmessage: %v", er.age, er.what)
}

// Person struct with age validation
type Person struct {
	age int
}

func (per *Person) getAge() (int, error) {
	if per.age < 0 || per.age > 100 {
		return per.age, &AgeError{per.age, "age can't be <0 or >100"}
	}
	return per.age, nil
}

// Sqrt function with error handling
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return x, nil
}

// Area function with error handling
func Area(x float64) (float64, error) {
	if x < 0 {
		return 0, SquareNegativeSide(x)
	}
	return x * x, nil
}

func main() {
	fmt.Println(Sqrt(2)) // Output: 2 <nil>
	fmt.Println(Sqrt(-3)) // Output: 0 cannot Sqrt negative number: -3

	p := Person{25}
	fmt.Println(Area(-9)) // Output: 0 error, side can't be negative! You input: -9
	fmt.Println(p.getAge()) // Output: 25 <nil>
}
```
---
## **🔹 Key Takeaways**
1️⃣ **Always initialize custom error types before returning them**  
   - **Wrong**: `e = n`
   - **Right**: `return 0, ErrNegativeSqrt(x)`

2️⃣ **Avoid unnecessary variable assignments**  
   - Define error messages in the `Error()` method **without using named return values**.

3️⃣ **Use struct-based errors when additional data is needed**  
   - Implement the `error` interface using a **pointer receiver** if you store additional fields.

4️⃣ **Check and return errors properly**  
   - Example:
   ```go
   if x < 0 {
       return 0, SquareNegativeSide(x)
   }
   ```

🚀 **After applying these fixes, the program works correctly and follows Go best practices!**