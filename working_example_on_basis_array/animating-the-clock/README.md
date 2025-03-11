# Digital Clock in Go Using ASCII Art  

## **Overview**  
This program creates a **live digital clock** in the terminal using ASCII-style characters. It updates every second to reflect the current time.  

## **Algorithm Explanation**  

### **1. Define `placeholder` for ASCII Digits**  
Each digit (0-9) and the colon (`:`) is represented as a **5-row ASCII art** using a `placeholder` type:  
```go
type placeholder [5]string
```
Example for `0`:  
```go
zero := placeholder{
    "███",
    "█ █",
    "█ █",
    "█ █",
    "███",
}
```
Similar representations are created for digits `0-9` and `:`.

### **2. Store All Digits in an Array**  
A fixed-size array stores all digit placeholders:  
```go
digits := [...]placeholder{zero, one, two, ..., nine}
```

### **3. Retrieve Current Time**  
The program gets the system's current hour, minute, and second:  
```go
now := time.Now()
hour, min, sec := now.Hour(), now.Minute(), now.Second()
```

### **4. Build the Clock Representation**  
The digits from the current time are used to construct an ASCII representation:  
```go
clock := [...]placeholder{
    digits[hour/10], digits[hour%10],
    colon,
    digits[min/10], digits[min%10],
    colon,
    digits[sec/10], digits[sec%10],
}
```
- Extracts individual digits from `hour`, `minute`, and `second`.  
- Inserts colons (`:`) between them.  

### **5. Print the Clock in Rows**  
The ASCII characters are printed row by row:  
```go
for line := range clock[0] {
    for digit := range clock {
        fmt.Print(clock[digit][line], "  ")
    }
    fmt.Println()
}
```
This loops through each **row** of the placeholders, printing each digit side by side.

### **6. Refresh Every Second**  
The screen is **cleared** and updated every second:  
```go
screen.Clear()
screen.MoveTopLeft()
time.Sleep(time.Second)
```
This ensures smooth real-time updates.

---

## **Example Output**
```
███  ███       ███  ███       ███  ███  
  █  █ █   ░   █    █     ░     █  █  
███  █ █       ███  ███       ███  ███  
█    █ █   ░     █    █   ░     █    █  
███  ███       ███  ███       ███  ███  
```
This represents a live clock displaying the current time.
