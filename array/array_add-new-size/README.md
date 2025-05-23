# **Array Modification in Go – Adding New Elements**

## **Overview**  
This program demonstrates the process of copying elements from one array to another, modifying them, and adding a new element.

---

## **Algorithm Explanation**  

### **1. Initialize the `prev` Array**
```go
prev := [3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}
```
- The `prev` array contains three books from the previous year.

### **2. Declare the `books` Array**  
```go
var books [4]string
```
- The `books` array is declared with a size of four, allowing space for one additional book.

### **3. Copy and Modify Elements**  
```go
for i, b := range prev {
    books[i] += b + " 2nd Ed."
}
```
- Each book title from `prev` is copied into `books`, adding the suffix `" 2nd Ed."`.

### **4. Add a New Book**  
```go
books[3] = "Awesomeness"
```
- A new book, `"Awesomeness"`, is added to `books` in the fourth index.

### **5. Print the Arrays**  
```go
fmt.Printf("last year:\n%#v\n", prev)
fmt.Printf("\nthis year:\n%#v\n", books)
```
- The contents of both arrays are printed to show the results.

---

## **Example Output**  
```
last year:
[3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

this year:
[4]string{"Kafka's Revenge 2nd Ed.", "Stay Golden 2nd Ed.", "Everythingship 2nd Ed.", "Awesomeness"}
```