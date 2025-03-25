This Go program contains a function named `CheckCondition`, which simulates an age verification system for a pub. Below is a breakdown of its functionality and logic.

---

## **📌 Functionality Overview**
The function checks whether customers are allowed to enter the pub based on their age. The **minimum age requirement is 18**. Three customers with different ages are checked against this condition.

---

## **🔍 Key Components**
### **1️⃣ Declaring Variables**
```go
allowed_age := 18

customer1 := 15
customer2 := 29
customer3 := 18
```
- `allowed_age`: The required age for entry (18 years old).
- `customer1`, `customer2`, and `customer3` are sample customers with ages **15, 29, and 18**.

---

### **2️⃣ Checking Conditions Using `if-else`**
The program checks each customer's age using **conditional statements**.

#### **🔹 Case 1: Customer is older than 18**
```go
if customer1 > allowed_age {
	fmt.Println(fmt.Sprintf("%v", customer1) + " You can enter, have fun!")
}
```
- If the customer's age **is greater than** `allowed_age` (18), they are allowed to enter.
- Example output for `customer2 = 29`:
  ```
  29 You can enter, have fun!
  ```

#### **🔹 Case 2: Customer is exactly 18**
```go
else if customer1 == allowed_age {
	fmt.Println("Oh, yo're just 18. But you can enter, have fun!")
}
```
- If the customer's age **is exactly 18**, they are still allowed to enter, but a different message is printed.
- Example output for `customer3 = 18`:
  ```
  Oh, yo're just 18. But you can enter, have fun!
  ```

#### **🔹 Case 3: Customer is younger than 18**
```go
else {
	fmt.Println(fmt.Sprintf("%v", customer1) + " You can't enter, sorry for that:(")
}
```
- If the customer **is younger than 18**, they are denied entry.
- Example output for `customer1 = 15`:
  ```
  15 You can't enter, sorry for that :(
  ```

---

## **🛠️ Code Optimization Suggestion**
Instead of repeating the same logic for three different customers, you can use a **loop**:
```go
func CheckCondition() {
	allowed_age := 18
	customers := []int{15, 29, 18}

	for _, age := range customers {
		if age > allowed_age {
			fmt.Println(fmt.Sprintf("%v", age) + " You can enter, have fun!")
		} else if age == allowed_age {
			fmt.Println("Oh, you're just 18. But you can enter, have fun!")
		} else {
			fmt.Println(fmt.Sprintf("%v", age) + " You can't enter, sorry for that :(")
		}
	}
}
```
This makes the code cleaner and easier to maintain. 🚀

Would you like additional modifications or explanations? 😊