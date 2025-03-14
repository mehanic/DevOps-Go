This Go program demonstrates how to **write data to a CSV file** and then **read it back**, printing the content to the console.

---

## ✅ **Step-by-step explanation:**

### 🎯 1. **Data to be Written to CSV:**
```go
dictionaries := []map[string]string{
	{"age": "30", "name": "John", "last_name": "Doe"},
	{"age": "30", "name": "Jane", "last_name": "Doe"},
}
```
- `dictionaries` is a **slice of maps**.
- Each map represents a **record (row)** with key-value pairs (the column name as the key and the value as the data).

---

### 🎯 2. **Open CSV File for Writing:**
```go
csvFile, err := os.Create("my.csv")
if err != nil {
	fmt.Println("Error creating file:", err)
	return
}
defer csvFile.Close()
```
- `os.Create("my.csv")` creates a new CSV file.
- `defer csvFile.Close()` ensures the file is closed after the function ends to avoid resource leaks.

---

### 🎯 3. **CSV Writer Initialization:**
```go
writer := csv.NewWriter(csvFile)
defer writer.Flush()
```
- `csv.NewWriter()` creates a **CSV writer object**.
- `defer writer.Flush()` ensures the writer writes any buffered data to the file before closing.

---

### 🎯 4. **Write CSV Header:**
```go
headers := []string{"age", "name", "last_name"}
if err := writer.Write(headers); err != nil {
	fmt.Println("Error writing header:", err)
	return
}
```
- The header defines the **column names** in the CSV file.

---

### 🎯 5. **Write the Records (Rows):**
```go
for _, record := range dictionaries {
	row := []string{
		record["age"],
		record["name"],
		record["last_name"],
	}
	if err := writer.Write(row); err != nil {
		fmt.Println("Error writing record:", err)
		return
	}
}
```
- The loop iterates through each **map (record)** in `dictionaries`.
- A **slice `row`** is created, extracting the values from the map.
- `writer.Write()` writes each row to the CSV file.

---

### 🎯 6. **Open CSV File for Reading:**
```go
csvFile, err = os.Open("my.csv")
if err != nil {
	fmt.Println("Error opening file:", err)
	return
}
defer csvFile.Close()
```
- Reopen the CSV file for reading.

---

### 🎯 7. **CSV Reader Initialization:**
```go
reader := csv.NewReader(csvFile)
```
- `csv.NewReader()` creates a **CSV reader object**.

---

### 🎯 8. **Read All Records:**
```go
records, err := reader.ReadAll()
if err != nil {
	fmt.Println("Error reading records:", err)
	return
}
```
- `reader.ReadAll()` reads all rows from the CSV file and stores them in the `records` variable.

---

### 🎯 9. **Print the Records:**
```go
for _, record := range records {
	fmt.Println(record)
}
```
- The loop prints each row from the `records` slice.

---

## ✅ **Final CSV Content (`my.csv`):**
```
age,name,last_name
30,John,Doe
30,Jane,Doe
```

---

## ✅ **Final Output in the Terminal:**
```
[age name last_name]
[30 John Doe]
[30 Jane Doe]
```

---

## 🔥 **Key Concepts Explained:**
| Concept               | Explanation                                      |
|-----------------|--------------------------------------------------|
| Slice of Maps     | Each map represents a CSV row with key-value pairs. |
| CSV Writer          | Writes data to the CSV file. |
| CSV Reader          | Reads data from the CSV file. |
| Error Handling     | Proper handling of file and CSV writing errors. |

---

Let me know if you'd like more examples or enhancements! 😊