This Go program demonstrates **decoding GOB-encoded data into different structures** using the `encoding/gob` package. Let's break it down:

---

## ✅ **What is GOB?**
- **GOB** is Go's **binary serialization format**, similar to JSON or XML but optimized for Go data types.
- It is efficient and can handle complex Go data structures like maps and slices.

---

## 🛑 **Structure Definition:**
```go
type Character struct {
	Name        string `gob:"name"`
	Surname     string `gob:"surname"`
	Job         string `gob:"job,omitempty"`
	YearOfBirth int    `gob:"year_of_birth,omitempty"`
}
```
- This is the **main struct (`Character`) that is being encoded and decoded.**
- The **struct tags (`gob:"name"`)** map the Go fields to GOB encoding.
- `omitempty` means **if the field is empty, it will be skipped during encoding.**

---

## 📝 **Encoded Data (GOB format):**
```go
data := []byte("D\xff\x81\x03\x01\x01\tCharacter" +
    "\x01\xff\x82\x00\x01\x04\x01\x04Name" +
    "\x01\f\x00\x01\aSurname\x01\f\x00\x01\x03" +
    "Job\x01\f\x00\x01\vYearOfBirth\x01\x04\x00" +
    "\x00\x00*\xff\x82\x01\x06Albert\x01\bWilmarth" +
    "\x01\x13assistant professor\x00")
```
- This is the **binary GOB-encoded data**, representing an instance of the `Character` struct:
  - Name: `"Albert"`
  - Surname: `"Wilmarth"`
  - Job: `"assistant professor"`
  - YearOfBirth: **(missing due to `omitempty`)**

---

## 🔎 **The `runDecode()` Function:**
```go
func runDecode(data []byte, v interface{}) {
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(v); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", v)
}
```
- It takes **GOB data (`data`)** and an **empty struct (`v`)** to decode into.
- It uses the `gob.NewDecoder()` to decode the binary data into the provided struct.
- If decoding fails, it logs an error and stops the program.

---

## 🎯 **Different Decoding Scenarios:**
### 1️⃣ **Exact Match (Decoding into the `Character` struct):**
```go
runDecode(data, new(Character))
```
- This works perfectly because the GOB data exactly matches the `Character` struct.
- Output:
```
{Name:Albert Surname:Wilmarth Job:assistant professor YearOfBirth:0}
```

---

### 2️⃣ **Changing the Field Order:**
```go
runDecode(data, new(struct {
	YearOfBirth int    `gob:"year_of_birth,omitempty"`
	Surname     string `gob:"surname"`
	Name        string `gob:"name"`
	Job         string `gob:"job,omitempty"`
}))
```
- The **order of fields doesn't matter**, as long as the field names and types match.
- Output:
```
{YearOfBirth:0 Surname:Wilmarth Name:Albert Job:assistant professor}
```

---

### 3️⃣ **Decoding Only One Field:**
```go
runDecode(data, new(struct {
	Name string `gob:"name"`
}))
```
- This works because **GOB allows partial decoding**.
- The missing fields are simply ignored.
- Output:
```
{Name:Albert}
```

---

### 4️⃣ **Decoding with Extra Unmatched Fields:**
```go
runDecode(data, new(struct {
	Name        string `gob:"name"`
	Surname     string `gob:"surname"`
	Country     string `gob:"country"` // Extra field not in GOB data
	Job         string `gob:"job,omitempty"`
	YearOfBirth int    `gob:"year_of_birth,omitempty"`
}))
```
- GOB **ignores unknown fields** (`Country`), so decoding still works.
- Output:
```
{Name:Albert Surname:Wilmarth Country: Job:assistant professor YearOfBirth:0}
```

---

### 5️⃣ **Decoding with Mismatched Field Type:**
```go
runDecode(data, new(struct {
	Name []byte `gob:"name"` // Expecting []byte instead of string
}))
```
- This **fails** because the GOB data expects a `string`, not a `[]byte`.
- Error:
```
gob: type mismatch: string vs. []uint8
```

---

## ✅ **Key Takeaways:**
| Scenario               | Works? | Reason |
|----------------|---------|------------------------|
| Exact struct match  | ✔️ | All fields match |
| Field order changed | ✔️ | GOB ignores field order |
| Partial decoding | ✔️ | GOB allows missing fields |
| Extra fields in target | ✔️ | GOB ignores unknown fields |
| Mismatched field types | ❌ | Type mismatch |
