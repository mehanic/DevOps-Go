This Go program contains multiple functions that demonstrate **if-else statements**, **logical conditions**, **command-line argument handling (`os.Args`)**, and **leap year calculations**. Below is an explanation of each function.

---

## **1. `main()` - Checking Even/Odd and Divisibility by 8**
```go
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
		return
	}

	if n%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8\n", n)
	} else if n%2 == 0 {
		fmt.Printf("%d is an even number\n", n)
	} else {
		fmt.Printf("%d is an odd number\n", n)
	}
}
```

### **How It Works:**
1. **Command-Line Argument Check**:  
   - If no number is provided, it prints `"Pick a number"` and exits.
2. **Convert Input to Integer**:  
   - Uses `strconv.Atoi()` to convert the string input into an integer.
   - If conversion fails, prints an error message and exits.
3. **Even or Odd Check**:
   - If the number is divisible by 8 (`n % 8 == 0`), it prints that the number is **even and divisible by 8**.
   - If it's even (`n % 2 == 0` but not divisible by 8), it prints **even number**.
   - Otherwise, it's **odd**.

### **Example Runs:**
```
$ go run main.go 16
16 is an even number and it's divisible by 8

$ go run main.go 10
10 is an even number

$ go run main.go 7
7 is an odd number

$ go run main.go abc
"abc" is not a number
```

---

## **2. `main1()` - Checking Leap Year (Verbose)**
```go
func main1() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	// Verbose way to determine if it's a leap year
	var leap bool
	if year%400 == 0 {
		leap = true
	} else if year%100 == 0 {
		leap = false
	} else if year%4 == 0 {
		leap = true
	}

	if leap {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
```

### **Leap Year Rules:**
- A year is a **leap year** if:
  - It is divisible by **400** → `year % 400 == 0`
  - OR it is divisible by **4** AND **not divisible by 100** → `year % 4 == 0 && year % 100 != 0`

### **How It Works:**
1. **Gets Year Input**:  
   - If no input is provided, it asks for one.
   - Converts input to an integer.
2. **Determines If It's a Leap Year**:
   - If divisible by **400**, it's a leap year.
   - If divisible by **100**, it's **not** a leap year.
   - If divisible by **4**, it's a leap year.
3. **Prints the Result**.

### **Example Runs:**
```
$ go run main.go 2000
2000 is a leap year.

$ go run main.go 1900
1900 is not a leap year.

$ go run main.go 2024
2024 is a leap year.

$ go run main.go abc
"abc" is not a valid year.
```

---

## **3. `main3()` - Optimized Leap Year Check**
```go
func main3() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
```
### **Differences from `main1()`**
- This uses a **single if condition** instead of multiple `if-else` blocks.
- The condition `year%4 == 0 && (year%100 != 0 || year%400 == 0)` checks all leap year rules **at once**.

### **Same Example Runs as `main1()`** but optimized.

---

## **4. `main4()` - Checking Month Days**
```go
func main4() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	// Get the current year and determine if it's a leap year
	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	// Default to 28 (for February in non-leap years)
	days := 28

	month := os.Args[1]

	// Convert to lowercase for case-insensitive comparison
	if m := strings.ToLower(month); m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		days = 30
	} else if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		days = 31
	} else if m == "february" {
		if leap {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
```

### **How It Works:**
1. **Gets the Month Name from User**.
2. **Determines the Current Year's Leap Year Status**.
3. **Checks the Number of Days for Each Month**:
   - **30 Days**: April, June, September, November.
   - **31 Days**: January, March, May, July, August, October, December.
   - **February**: 
     - **28 days** in regular years.
     - **29 days** in leap years.
4. **Prints the Number of Days**.

### **Example Runs:**
```
$ go run main.go March
"March" has 31 days.

$ go run main.go april
"april" has 30 days.

$ go run main.go February
"February" has 28 days.   # (Non-leap year)

$ go run main.go February
"February" has 29 days.   # (Leap year - depends on current year)

$ go run main.go abc
"abc" is not a month.
```

---

## **Summary of Rules**
1. **Even/Odd & Divisibility**:
   - `% 8 == 0` → Even and divisible by 8.
   - `% 2 == 0` → Even.
   - Else → Odd.

2. **Leap Year Check**:
   - `% 400 == 0` → Leap year.
   - `% 100 == 0` → NOT a leap year.
   - `% 4 == 0` → Leap year.
   - Optimized version: `year%4 == 0 && (year%100 != 0 || year%400 == 0)`

3. **Month Days**:
   - **31 Days**: Jan, Mar, May, Jul, Aug, Oct, Dec.
   - **30 Days**: Apr, Jun, Sep, Nov.
   - **Feb**: 28 or 29 (Leap year).

---

This covers everything! 🚀 Let me know if you have any questions. 😃