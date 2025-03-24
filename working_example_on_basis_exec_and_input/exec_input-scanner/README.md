## **Program Explanation**
This Go program **asks the user for their first and last name, formats it properly, and then prints a greeting**. The user can enter `'q'` at any time to quit.

---

## **Step-by-Step Breakdown**

### **1. Function: `getFormattedName`**
```go
func getFormattedName(firstName, lastName string) string {
	fullName := firstName + " " + lastName
	return strings.Title(fullName)
}
```
- **Concatenates** `firstName` and `lastName` into `fullName`.
- **Uses `strings.Title`** to capitalize the first letter of each word.
- **Returns the formatted name**.

🔹 **Example:**  
```go
getFormattedName("max", "ned") → "Max Ned"
```

---

### **2. Setting Up the Scanner**
```go
scanner := bufio.NewScanner(os.Stdin)
```
- **Creates a scanner** to read user input.

---

### **3. Infinite Loop (`for {}`)**
```go
for {
	fmt.Println("\nPlease tell me your name: ")
	fmt.Println("(enter 'q' at any time to quit)")
```
- The program **keeps asking for names indefinitely** until the user enters `'q'`.

---

### **4. Getting the First Name**
```go
fmt.Print("First name: ")
scanner.Scan()
fName := scanner.Text()
if strings.ToLower(fName) == "q" {
	break
}
```
- **Reads user input for first name**.
- **Checks if the user entered `'q'` (case-insensitive)** → If so, exits the loop.

---

### **5. Getting the Last Name**
```go
fmt.Print("Last name: ")
scanner.Scan()
lName := scanner.Text()
if strings.ToLower(lName) == "q" {
	break
}
```
- **Reads user input for last name**.
- **Exits if the user entered `'q'`**.

---

### **6. Formatting and Printing the Name**
```go
formattedName := getFormattedName(fName, lName)
fmt.Printf("\nHello, %s!\n", formattedName)
```
- Calls `getFormattedName` to **capitalize the first letter of each word**.
- Prints a **greeting with the formatted name**.

---

## **Example Runs**

### **Run 1 (Normal Case)**
#### **Input:**
```
Please tell me your name: 
(enter 'q' at any time to quit)
First name: max
Last name: ned
```
#### **Output:**
```
Hello, Max Ned!
```

---

### **Run 2 (Quitting Early)**
#### **Input:**
```
Please tell me your name: 
(enter 'q' at any time to quit)
First name: q
```
#### **Output:**
```
(nothing, program exits)
```
- The user enters `'q'`, so the program **terminates immediately**.

---

### **Run 3 (Quitting After First Name)**
#### **Input:**
```
Please tell me your name: 
(enter 'q' at any time to quit)
First name: John
Last name: q
```
#### **Output:**
```
(nothing, program exits)
```
- The program quits **before printing the greeting**.

---

## **Key Features**
✅ **Loops until the user decides to quit**  
✅ **Formats names properly (capitalizes the first letter)**  
✅ **Allows quitting at any point (`q` to exit)**  
✅ **Handles different cases (e.g., lowercase input, quitting mid-way)**  

Would you like any modifications, such as adding middle names or handling empty inputs? 🚀