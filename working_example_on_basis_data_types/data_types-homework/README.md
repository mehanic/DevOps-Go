This Go program contains multiple tasks that implement simple algorithms and mathematical operations. Each function handles a specific task, and the `main` function calls each of these tasks with various parameters. Let's break down the code and explain the logic of each function.

### **Task Descriptions and Explanation:**

#### **Task 1 (`tastk_1`) - FizzBuzz Algorithm**
```go
func tastk_1(n int) {
	if n%5 == 0 && n%3 == 0 {
		fmt.Println("FizzBuzz")
	} else if n%3 == 0 {
		fmt.Println("Fizz")
	} else if n%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(n)
	}
}
```
- This is a classic **FizzBuzz** program where:
  - If a number is divisible by both 3 and 5, print "FizzBuzz".
  - If a number is divisible by 3, print "Fizz".
  - If a number is divisible by 5, print "Buzz".
  - Otherwise, print the number itself.

#### **Task 2 (`tastk_2`) - Palindrome Check**
```go
func tastk_2(n int) {
	a, b, c := n/100, n%100/10, n%10
	d := c*100 + b*10 + a
	if d == n {
		fmt.Println("Palidrom")
	} else {
		fmt.Println("It isn't Palidrom")
	}
}
```
- This function checks if a 3-digit number is a **palindrome**:
  - A **palindrome** is a number that reads the same forwards and backwards. Example: `121` is a palindrome.
  - The code extracts the hundreds, tens, and ones place from the number and checks if reversing the digits produces the original number.

#### **Task 3 (`task_3`) - Swap Variables**
```go
func task_3(a, b, c int) {
	fmt.Printf("a =: %v b =: %v c =: %v\n", a, b, c)
	a, b, c = c, a, b
	fmt.Printf("a =: %v b =: %v c =: %v\n", a, b, c)
}
```
- This function swaps the values of three variables (`a`, `b`, and `c`), and prints the values before and after swapping. It demonstrates Go's multiple assignment feature.

#### **Task 4 (`task_4`) - Find the Smallest Number**
```go
func task_4(x, y, z int) {
	if x <= y && x <= z {
		fmt.Println(x)
	} else if y <= x && y <= z {
		fmt.Println(y)
	} else if z <= x && z <= y {
		fmt.Println(z)
	}
}
```
- This function finds the **smallest number** from three input integers and prints it.

#### **Task 5 (`task_5`) - Find the Largest Number**
```go
func task_5(x, y, z int) {
	if x >= y && x >= z {
		fmt.Println(x)
	} else if y >= x && y >= z {
		fmt.Println(y)
	} else if z >= x && z >= y {
		fmt.Println(z)
	}
}
```
- Similar to Task 4, this function finds the **largest number** from three input integers.

#### **Task 6 (`task_6`) - Check if Number is Even or Odd**
```go
func task_6(n int) {
	if n%2 == 0 {
		fmt.Println("even \"Juft\"")
	} else {
		fmt.Println("odd \"Toq\"")
	}
}
```
- This function checks if a number is **even** or **odd**:
  - If the number is divisible by 2, it prints "even" (or "Juft" in Uzbek).
  - Otherwise, it prints "odd" (or "Toq" in Uzbek).

#### **Task 7 (`task_7`) - Leap Year Check**
```go
func task_7(year int) {
	if year%4 == 0 {
		fmt.Println("It year")
	} else {
		fmt.Println("It is not a leap year")
	}
}
```
- This function checks if a given year is a **leap year**. In this code, the leap year condition is simplified to checking if the year is divisible by 4.

#### **Task 8 (`task_8`) - Calculate Average**
```go
func task_8(A, B, C float32) {
	fmt.Println((A + B + C) / 3)
}
```
- This function calculates the **average** of three numbers (`A`, `B`, `C`).

#### **Task 9 (`task_9`) - Simple Calculator**
```go
func task_9(a, b float32, operator string) {
	if operator == "+" {
		fmt.Printf("%v + %v = %v\n", a, b, a+b)
	} else if operator == "-" {
		fmt.Printf("%v - %v = %v\n", a, b, a-b)
	} else if operator == "*" {
		fmt.Printf("%v * %v = %v\n", a, b, a*b)
	} else if operator == "/" && b != 0 {
		fmt.Printf("%v / %v = %v\n", a, b, a/b)
	} else if operator == "/" && b == 0 {
		fmt.Printf("%v sonini %v soniga Bo'lib bo'lmaydi\n", a, b)
	} else {
		fmt.Printf("\"%s\" Bunaqa matematik operator yo'q\n", operator)
	}
}
```
- This function implements a simple **calculator** that can perform addition, subtraction, multiplication, and division based on the input operator.
  - It also handles division by zero and unknown operators.

#### **Task 10 (`task_10`) - Financial Calculation**
```go
func task_10() {
	var (
		investorPercentage = 70
		companyPercentage  = 30
		malibuPrice        = 700_000
		insuranePrice      = 100_000
		orderDay           = 2
		totalSum           = malibuPrice*orderDay + insuranePrice*orderDay
		investor           = totalSum * investorPercentage / 100
		company            = totalSum*companyPercentage/100 - insuranePrice*orderDay
		insurance          = insuranePrice * orderDay
	)

	fmt.Println("Total Summa: ", totalSum)
	fmt.Println("Investor: ", investor)
	fmt.Println("Company: ", company)
	fmt.Println("Insurance: ", insurance)
}
```
- This function calculates the **financial distribution** for an investor, company, and insurance.
  - It calculates the total price, investor share, company share, and insurance cost based on the given percentages.

#### **Task 11 (`task_11`) - Square and Perimeter**
```go
func task_11(a int) {
	fmt.Println("S = ", a*a, "\nP = ", 4*a)
}
```
- This function calculates the **area** (`S`) and **perimeter** (`P`) of a square given its side length `a`.

#### **Task 12 (`task_12`) - Temperature Conversion**
```go
func task_12(selse float32) {
	fmt.Println("Kelven =", selse*273.15)
	fmt.Println("Fahrenheit = ", selse*1.8+32)
}
```
- This function converts a temperature in **Celsius** to **Kelvin** and **Fahrenheit**.

#### **Task 13 (`task_13`) - Day of the Week**
```go
func task_13(day int) {
	if day == 1 {
		fmt.Println("Dushanba")
	} else if day == 2 {
		fmt.Println("Seshanba")
	} else if day == 3 {
		fmt.Println("Chorshanba")
	} else if day == 4 {
		fmt.Println("Payshanba")
	} else if day == 5 {
		fmt.Println("Juma")
	} else if day == 6 {
		fmt.Println("Shanba")
	} else if day == 7 {
		fmt.Println("Yakshanba")
	}
}
```
- This function converts an integer representing a day of the week into the corresponding weekday name in Uzbek.
  - 1 = Monday, 2 = Tuesday, ..., 7 = Sunday.

#### **Task 14 (`task_14`) - Swap Digits**
```go
func task_14(n int) {
	fmt.Println(n)
	a, b := n%100/10, n%10
	a, b = b, a
	fmt.Println(n - n%100 + a*10 + b)
}
```
- This function swaps the **last two digits** of a given number `n` and prints the result.

---

### **Conclusion:**
Each function in this program is designed to implement a simple algorithm or perform a basic calculation. The tasks cover a wide range of topics like FizzBuzz, leap year checks, temperature conversions, and basic mathematical operations. These tasks are useful for understanding core programming concepts like conditionals, loops, data types, and functions in Go.