### **📌 Explanation of the Rating Scale Conversion in Go**  

This Go program takes a user rating between **0 and 5** and converts it to a scale of **0 to 100** by multiplying it by **20**.

---

## **🔹 Code Breakdown**  

### **1️⃣ Read User Input**
```go
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a rating between 0 to 5")
	input, _ := reader.ReadString('\n')
```
✅ `bufio.NewReader(os.Stdin)` is used to read input from the console.  
✅ `ReadString('\n')` captures the user input until the user presses **Enter**.  
✅ `_` (underscore) ignores any errors that might occur while reading input.  

---

### **2️⃣ Display the Raw Input**
```go
	fmt.Println("Your Input is ", input)
```
✅ This prints the user's input as received, including any spaces or newline characters.

---

### **3️⃣ Convert the Rating to a Float**
```go
	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
```
✅ `strings.TrimSpace(input)` removes any leading/trailing spaces or newline characters.  
✅ `strconv.ParseFloat()` converts the input string to a `float64`.  
✅ If the input is invalid (e.g., letters instead of numbers), `err` will **not be `nil`**.

---

### **4️⃣ Scale the Rating to 0–100**
```go
	newRating = newRating * 20
```
✅ Since the rating was originally between **0 and 5**, multiplying it by **20** scales it to **0 to 100**.

---

### **5️⃣ Handle Errors and Display the Converted Rating**
```go
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New rating is ", newRating)
	}
```
✅ If an error occurs (`err != nil`), it prints the error message.  
✅ Otherwise, it prints the **scaled rating**.

---

## **🔹 Example Run**
### **Valid Input:**
```
Enter a rating between 0 to 5
3
Your Input is  3

changing scale to 0 to 100
New rating is  60
```
### **Invalid Input (e.g., a letter instead of a number)**
```
Enter a rating between 0 to 5
abc
Your Input is  abc

changing scale to 0 to 100
strconv.ParseFloat: parsing "abc": invalid syntax
```
👉 The program catches the error and prints an appropriate message.

---

## **🔹 Key Takeaways**
✅ `bufio.NewReader(os.Stdin)` reads user input.  
✅ `strings.TrimSpace()` removes unwanted spaces/newlines.  
✅ `strconv.ParseFloat()` converts strings to floating-point numbers.  
✅ A simple `if-else` handles error checking.  
✅ The rating is **scaled** by multiplying it by **20**.  

---

Would you like **improvements**, such as handling edge cases (e.g., numbers outside 0–5) or rounding the output? 😊