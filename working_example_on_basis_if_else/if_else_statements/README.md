### **📌 Explanation of `if`, `else if`, and Conditionals in Go**  

This program demonstrates **conditional statements (`if`, `else if`, `else`)**, **nested conditions (`ifception`)**, and **comparison operators** in Go.

---

## **🔹 Code Breakdown**  

### **1️⃣ `if, else` Statement**
```go
	number1 := 10
	number2 := 20

	if number1 > number2 {
		fmt.Println("10 is bigger than 20")
	} else {
		fmt.Println("No, 10 is not bigger than 20.")
	}
```
✅ `if number1 > number2` checks if `10 > 20`.  
✅ Since this condition is **false**, the `else` block runs.  
✅ **Output:**
```
No, 10 is not bigger than 20.
```

---

### **2️⃣ `if, else if, else` Statement**
```go
	if 9%2 == 0 {
		fmt.Println("9 is a even number")
	} else if 9 > 10 {
		fmt.Println("9 is also bigger than 10!")
	} else {
		fmt.Println("No, 9 isn't a even number nor is it bigger than 10.")
	}
```
✅ `if 9%2 == 0` checks if `9` is even (false).  
✅ `else if 9 > 10` checks if `9` is greater than `10` (false).  
✅ Since both are false, the `else` block runs.  
✅ **Output:**
```
No, 9 isn't a even number nor is it bigger than 10.
```

---

### **3️⃣ Nested `if` (Ifception 😆)**
```go
	if number1 > 9 {
		if number2 > 19 {
			fmt.Println("Both conditions are true")
		} else {
			fmt.Println("Only the first condition is true")
		}
	} else {
		fmt.Println("The first condition is false")
	}
```
✅ `if number1 > 9` checks if `10 > 9` (**true**), so it enters the first block.  
✅ Inside it, `if number2 > 19` checks if `20 > 19` (**true**), so it prints:  
```
Both conditions are true
```

---

### **4️⃣ Comparison Operators**
```go
	fmt.Println(number1 < number2)  // true
	fmt.Println(number1 > number2)  // false
	fmt.Println(number1 <= number2) // true
	fmt.Println(number1 >= number2) // false
	fmt.Println(number1 == number2) // false
	fmt.Println(number1 != number2) // true
```
✅ These statements **compare `number1 (10)` and `number2 (20)`** and print **Boolean values (`true/false`)**.

**Output:**
```
true   // 10 < 20
false  // 10 > 20
true   // 10 <= 20
false  // 10 >= 20
false  // 10 == 20
true   // 10 != 20
```

---

## **🔹 Key Takeaways**
✅ `if` checks a condition; `else` runs if `if` is false.  
✅ `else if` allows checking multiple conditions.  
✅ **Nested `if`** statements allow deeper logical checks.  
✅ **Comparison operators (`==`, `!=`, `<`, `>`, `<=`, `>=`)** return `true` or `false`.  

Would you like an example with **logical operators (`&&`, `||`)** or **switch cases**? 😊