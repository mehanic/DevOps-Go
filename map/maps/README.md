This Go program demonstrates **working with maps** in different ways, including **nested maps**, **checking for key existence**, and **handling slices inside maps**. Let's break it down step by step.

---

## **1. Declaring Maps**
The program starts by defining several maps:

### **1️⃣ Simple Key-Value Map (Single Value)**
```go
phones := map[string]string{
    "bowen": "202-555-0179",
    "dulin": "03.37.77.63.06",
    "greco": "03489940240",
}
```
- Maps names (`string`) to **a single phone number** (`string`).

### **2️⃣ Boolean Map**
```go
products := map[int]bool{
    617841573: true,
    879401371: false,
    576872813: true,
}
```
- Maps **product IDs (`int`)** to **availability (`bool`)**.
- `true` → Available, `false` → Not available.

### **3️⃣ Map with Slices (Multiple Values per Key)**
```go
multiPhones := map[string][]string{
    "bowen": {"202-555-0179"},
    "dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
    "greco": {"03489940240", "03489900120"},
}
```
- **Maps names (`string`) to slices of phone numbers (`[]string`)**.
- Some people have **multiple phone numbers**.

### **4️⃣ Nested Maps (Map of Maps)**
```go
basket := map[int]map[int]int{
    100: {617841573: 4, 576872813: 2},
    101: {576872813: 5, 657473833: 20},
    102: {},
}
```
- **Maps customer IDs (`int`) to another map**.
- The inner map stores **product IDs (`int`)** and their **quantities (`int`)** in the shopping basket.

---

## **2. Fetching & Printing Values**

### **🔹 Print Dulin's Phone Number**
```go
who, phone := "dulin", "N/A"
if v, ok := phones[who]; ok {
    phone = v
}
fmt.Printf("%s's phone number: %s\n", who, phone)
```
- **Key existence check (`ok` idiom)**:  
  - If `"dulin"` exists in `phones`, assign its value to `phone`.
  - If not, use `"N/A"`.
- Prints:  
  ```
  dulin's phone number: 03.37.77.63.06
  ```

---

### **🔹 Check Product Availability**
```go
id, status := 879401371, "available"
if !products[id] {
    status = "not " + status
}
fmt.Printf("Product ID #%d is %s\n", id, status)
```
- If `products[id]` is `false`, change `status` to `"not available"`.
- Prints:
  ```
  Product ID #879401371 is not available
  ```

---

### **🔹 Get Greco's Second Phone Number**
```go
who, phone = "greco", "N/A"
if phones := multiPhones[who]; len(phones) >= 2 {
    phone = phones[1]
}
fmt.Printf("%s's 2nd phone number: %s\n", who, phone)
```
- **Retrieves a list of phone numbers from `multiPhones`**.
- If the list has **at least 2 numbers**, print the second one (`phones[1]`).
- Prints:
  ```
  greco's 2nd phone number: 03489900120
  ```

---

### **🔹 How Many Products is Customer #101 Buying?**
```go
cid, pid := 101, 576872813
fmt.Printf("Customer #%d is going to buy %d from Product ID #%d.\n", cid, basket[cid][pid], pid)
```
- Retrieves the **quantity** of `pid=576872813` for `cid=101` from `basket`.
- Prints:
  ```
  Customer #101 is going to buy 5 from Product ID #576872813.
  ```

---

## **Key Takeaways**
✅ **Maps store key-value pairs and can contain complex data like slices or other maps**.  
✅ **Checking if a key exists (`ok` idiom) prevents errors when accessing a map**.  
✅ **Nested maps allow modeling hierarchical data like shopping baskets**.  
✅ **Slices in maps help store multiple values per key** (e.g., multiple phone numbers per person).  

This program demonstrates **map manipulation, conditional lookups, and nested structures** in Go! 🚀


This Go program demonstrates how to **iterate over a map using a `for` loop with `range`**.

---

## **🔹 Code Breakdown**
### **1️⃣ Defining a Map**
```go
salaries := map[string]int{"Berkay": 10000, "John": 185000, "Kevin": 7200, "Son": 22000, "Pethy": 15000}
```
- **Declares a map** named `salaries`.
- **Keys** are `string` (employee names).
- **Values** are `int` (salary amounts in dollars).
- The map contains **five key-value pairs**.

---

### **2️⃣ Iterating Over the Map**
```go
for name, salary := range salaries {
    fmt.Printf("Salary of %v is %v $\n", name, salary)
}
```
- **`for ... range` loop** iterates over the `salaries` map.
- On each iteration:
  - `name` holds the **employee's name** (key).
  - `salary` holds the **employee's salary** (value).
- **`fmt.Printf` formats and prints each employee’s salary**.

---

## **🔹 Expected Output (Order May Vary)**
```bash
Salary of John is 185000 $
Salary of Berkay is 10000 $
Salary of Kevin is 7200 $
Salary of Son is 22000 $
Salary of Pethy is 15000 $
```
> **Note:** Map iteration order is **not guaranteed** in Go (maps are unordered collections).

---

## **🔹 Key Takeaways**
✅ **Maps store key-value pairs** in Go.  
✅ **`for ... range` loops allow easy iteration over maps**.  
✅ **Order of iteration is random** (not sorted).  
✅ **`fmt.Printf` is used for formatted output**.  

This function is useful for **displaying employee salaries, statistics, or reporting data from a map**. 🚀