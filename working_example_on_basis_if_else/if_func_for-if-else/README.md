Let's break down the code and explain the rules, step by step:

### **Code Overview:**
This Go program defines a function `isPhoneNumber()` to check if a given string is in the format of a phone number. The format it checks for is:
- 3 digits, followed by a hyphen (`-`), 3 more digits, another hyphen (`-`), and finally 4 digits (e.g., `415-555-4242`).

The program then tests the function with two inputs:
1. A valid phone number: `"415-555-4242"`.
2. A non-phone number string: `"Moshi moshi"`.

### **1. Importing Packages:**
```go
import (
	"fmt"
	"unicode"
)
```
- `fmt`: This package is used for formatted I/O, like printing to the console.
- `unicode`: This package provides functions to test characters for properties like whether they are digits. Specifically, the function `unicode.IsDigit()` is used to check if a character is a numeric digit.

### **2. `isPhoneNumber()` Function:**
The `isPhoneNumber()` function takes a string `text` as input and checks if it matches the expected phone number format.

#### **Step-by-step Explanation:**

1. **Check if the length of the string is 12 characters:**
   ```go
   if len(text) != 12 {
       return false // not phone number-sized
   }
   ```
   - The function first checks if the length of the string is 12 characters. 
   - If it's not, the function immediately returns `false` because a valid phone number must have exactly 12 characters in this format: `XXX-XXX-XXXX`.

2. **Check the first 3 characters for digits (area code):**
   ```go
   for i := 0; i < 3; i++ {
       if !unicode.IsDigit(rune(text[i])) {
           return false // not an area code
       }
   }
   ```
   - The loop checks the first three characters of the string (`text[0]`, `text[1]`, `text[2]`) to see if they are digits.
   - If any of the first three characters is not a digit, the function returns `false`.

3. **Check the fourth character for a hyphen:**
   ```go
   if text[3] != '-' {
       return false // does not have first hyphen
   }
   ```
   - The function checks if the fourth character is a hyphen (`-`).
   - If it's not a hyphen, it returns `false`.

4. **Check the next 3 characters for digits:**
   ```go
   for i := 4; i < 7; i++ {
       if !unicode.IsDigit(rune(text[i])) {
           return false // does not have first 3 digits
       }
   }
   ```
   - The loop checks characters 4 to 6 (`text[4]`, `text[5]`, `text[6]`) to ensure they are digits.
   - If any of these characters is not a digit, the function returns `false`.

5. **Check the eighth character for a hyphen:**
   ```go
   if text[7] != '-' {
       return false // does not have second hyphen
   }
   ```
   - The function checks if the eighth character is a hyphen (`-`).
   - If it's not, the function returns `false`.

6. **Check the last 4 characters for digits:**
   ```go
   for i := 8; i < 12; i++ {
       if !unicode.IsDigit(rune(text[i])) {
           return false // does not have last 4 digits
       }
   }
   ```
   - The loop checks the last four characters of the string (`text[8]`, `text[9]`, `text[10]`, `text[11]`) to ensure they are digits.
   - If any of these characters is not a digit, the function returns `false`.

7. **Return `true` if all checks pass:**
   ```go
   return true // text is a phone number
   ```
   - If all checks pass, the function returns `true`, meaning the input string is a valid phone number in the expected format.

### **3. `main()` Function:**
```go
func main() {
	// Test cases
	fmt.Println("415-555-4242 is a phone number:")
	fmt.Println(isPhoneNumber("415-555-4242"))
	fmt.Println("Moshi moshi is a phone number:")
	fmt.Println(isPhoneNumber("Moshi moshi"))
}
```

- **Test Case 1:** `"415-555-4242"` is a valid phone number format. The function `isPhoneNumber()` will check each part of this string:
  - The first three characters (`"415"`) are digits.
  - The fourth character is a hyphen.
  - The next three characters (`"555"`) are digits.
  - The eighth character is a hyphen.
  - The last four characters (`"4242"`) are digits.
  - Since all conditions are satisfied, the function returns `true`, and the program will print:
    ```
    415-555-4242 is a phone number:
    true
    ```

- **Test Case 2:** `"Moshi moshi"` is not a valid phone number. The string does not have the correct length (12 characters) or format. The function will return `false` because it doesn't match the expected phone number format:
  ```
  Moshi moshi is a phone number:
  false
  ```

### **Summary of the Rules:**
1. **Phone Number Format:**
   - The phone number must be exactly 12 characters long.
   - The format must be `XXX-XXX-XXXX` where `X` is a digit, and the hyphens (`-`) are in the correct places.
   
2. **Validation:**
   - The first three characters must be digits (the area code).
   - The fourth and eighth characters must be hyphens.
   - The middle three characters and last four characters must be digits.
   
3. **Return Values:**
   - The function returns `true` if the input string matches the format of a phone number.
   - It returns `false` if the input string does not meet the format requirements.

This approach ensures that the input string is validated strictly according to the format `XXX-XXX-XXXX`.