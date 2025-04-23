This Go program checks whether two numbers are **amicable numbers**. 

---

## **What are Amicable Numbers?**
Two numbers are considered **amicable** if the **sum of the proper divisors** of one number is **equal to the other number**, and vice versa.

For example:
- 220 and 284 are amicable because:
  - The sum of the divisors of 220: `1 + 2 + 4 + 5 + 10 + 11 + 20 + 22 + 44 + 55 + 110 = 284`
  - The sum of the divisors of 284: `1 + 2 + 4 + 71 + 142 = 220`

---

## ✅ **Code Explanation**

### 1. **Function `AmicableNumberCheck()`**
This function handles the main logic.

### 2. **Input Handling**
```go
var number1, number2 int
fmt.Scanln(&number1)
fmt.Scanln(&number2)
```
- The user is prompted to input two numbers.

---

### 3. **Calculating the Sum of Divisors**
```go
number1_dividors_sum := 0
number2_dividors_sum := 0
bigger := int(math.Max(float64(number1), float64(number2)))
```
- Two variables `number1_dividors_sum` and `number2_dividors_sum` store the sum of the divisors.
- The `bigger` variable is used to determine the larger number, to limit the loop to that value.

---

### 4. **Loop through Divisors**
```go
for i := 1; i < bigger; i++ {
	if (number1 > i) && (number1%i == 0) {
		number1_dividors_sum = number1_dividors_sum + i
	}

	if (number2 > i) && (number2%i == 0) {
		number2_dividors_sum = number2_dividors_sum + i
	}
}
```
- For each number, it checks for **divisibility** (i.e., `number1 % i == 0`).
- If `i` is a divisor, it adds `i` to the sum of divisors for that number.

---

### 5. **Checking Amicable Condition**
```go
if (number2 == number1_dividors_sum) && (number1 == number2_dividors_sum) {
	fmt.Printf("Hurrrrayy! Numbers %v and %v are amicable", number1, number2)
} else {
	fmt.Printf("Sorry. Sum of divisors in %v is %v while in %v is %v\n", 
		number1, number1_dividors_sum, number2, number2_dividors_sum)
}
```
- If the sum of the divisors of `number1` equals `number2`, and vice versa, the numbers are amicable.
- Otherwise, it prints the sums of the divisors for both numbers.

---

## 🎯 **Example Output**
```
We will check if the numbers you give are amicable or not!
Please enter the first number
220
Please enter the second number
284
Hurrrrayy! Numbers 220 and 284 are amicable
```

---

## ✅ **Summary of Rules**
1. Prompt the user for two numbers.
2. Calculate the **sum of divisors** for each number.
3. Check if **number1's divisor sum = number2** and vice versa.
4. Output the result.