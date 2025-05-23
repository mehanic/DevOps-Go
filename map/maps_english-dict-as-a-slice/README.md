This Go program is a simple **English-to-Turkish word translator** that uses **arrays and a loop** to find translations.

---

## **🔹 Code Breakdown**
### **1️⃣ Getting Command-Line Arguments**
```go
args := os.Args[1:]
if len(args) != 1 {
    fmt.Println("[english word] -> [turkish word]")
    return
}
query := args[0]
```
- Uses `os.Args[1:]` to get **command-line arguments** (ignoring the program name).
- If **no argument or more than one** is given, it prints a usage message and exits.
- Stores the input word in `query`.

---

### **2️⃣ Defining English and Turkish Word Lists**
```go
english := []string{"good", "great", "perfect"}
turkish := []string{"iyi", "harika", "mükemmel"}
```
- Two **parallel slices**:
  - `english` holds **English words**.
  - `turkish` holds their **Turkish translations**.

---

### **3️⃣ Searching for the Input Word**
```go
for i, w := range english {
    if query == w {
        fmt.Printf("%q means %q\n", w, turkish[i])
        return
    }
}
```
- Uses a **`for ... range` loop** to iterate through the `english` slice.
- If `query` **matches** an English word (`w`), it prints its **Turkish translation** using the same index `i`.
- **Time Complexity**: **O(n)** (linear search).

---

### **4️⃣ Handling Missing Words**
```go
fmt.Printf("%q not found\n", query)
```
- If no match is found, it prints that the word **is not in the dictionary**.

---

## **🔹 Example Usage**
### **✅ When Word Exists**
```bash
$ go run main.go good
"good" means "iyi"
```
### **❌ When Word Doesn't Exist**
```bash
$ go run main.go bad
"bad" not found
```
### **⚠️ When No Input Is Given**
```bash
$ go run main.go
[english word] -> [turkish word]
```

---

## **🔹 Key Takeaways**
✅ **Uses `os.Args` to get user input from the command line.**  
✅ **Stores English and Turkish words in slices.**  
✅ **Uses a `for ... range` loop to search for a word (O(n) complexity).**  
✅ **Handles errors for missing words or incorrect input.**  

### **🔹 Possible Improvements**
🔹 Use a **map** (`map[string]string`) instead of slices for **faster lookups (O(1))**.  
🔹 Support **more words** by expanding the dictionary.  

Would you like to see an optimized version using a map? 🚀