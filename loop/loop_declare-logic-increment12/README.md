### Explanation of Each Function:

This Go code contains multiple functions that demonstrate different types of loops for iterating through collections such as slices, arrays, maps, and strings. Let's break down each function:

---

### 1. **Iterating through a slice** (`main` function):
```go
func main() {
	mySlice := []string{"Prasetiyo", "Angga", "Diana"}

	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}
}
```
- **Purpose:** Iterates over a slice (`mySlice`) and prints each element.
- **How it works:**
  - A **`for`** loop is used with a traditional initialization (`i := 0`), condition (`i < len(mySlice)`), and increment (`i++`).
  - The loop iterates from `i = 0` to `i = len(mySlice)-1`, and during each iteration, it prints the element at the `i`-th index of the slice.

**Output:**
```
Prasetiyo
Angga
Diana
```

---

### 2. **Iterating through a string** (`main1` function):
```go
func main1() {
	myString := "Sirclo"
	for _, eachString := range myString {
		fmt.Println(string(eachString))
	}
}
```
- **Purpose:** Iterates through a string and prints each character.
- **How it works:**
  - A **`for`** loop with the **`range`** keyword is used.
  - The **`range`** keyword returns both the index and the character (`eachString`) in the string `myString`.
  - The underscore (`_`) is used to ignore the index value because only the character is needed.
  - Each character is printed after converting it into a string using `string(eachString)`.

**Output:**
```
S
i
r
c
l
o
```

---

### 3. **Iterating through an array** (`main2` function):
```go
func main2() {
	myArray := [5]string{"Prasetiyo", "Diana", "Angga", "Satria", "Ramdani"}
	for _, eachArray := range myArray {
		fmt.Println(eachArray)
	}
}
```
- **Purpose:** Iterates over an array and prints each element.
- **How it works:**
  - The **`range`** keyword is used to loop through the array `myArray`.
  - The loop returns both the index and the value of each element, but the index is ignored using the underscore (`_`).
  - Each value (`eachArray`) is printed in the loop.

**Output:**
```
Prasetiyo
Diana
Angga
Satria
Ramdani
```

---

### 4. **Iterating through a slice (alternative)** (`main3` function):
```go
func main3() {
	mySlice := []string{"Prasetiyo", "Diana", "Angga", "Satria", "Ramdani"}
	for _, eachSlice := range mySlice {
		fmt.Println(eachSlice)
	}
}
```
- **Purpose:** Iterates over a slice and prints each element.
- **How it works:**
  - Similar to the previous function `main2`, this function uses a **`range`** loop to iterate through the slice `mySlice`.
  - The loop ignores the index (using `_`) and directly prints the value of each element (`eachSlice`).

**Output:**
```
Prasetiyo
Diana
Angga
Satria
Ramdani
```

---

### 5. **Iterating through a map** (`main4` function):
```go
func main4() {
	myMap := map[string]string{
		"Name":  "Prasetiyo",
		"Hobby": "Coding",
	}

	for _, eachMap := range myMap {
		fmt.Println(eachMap)
	}
}
```
- **Purpose:** Iterates through a map and prints each value.
- **How it works:**
  - The **`range`** keyword is used to iterate over the map `myMap`.
  - The loop returns both the key and the value, but only the value (`eachMap`) is printed (key is ignored with `_`).

**Output:**
```
Prasetiyo
Coding
```
Note: The order of the map iteration is not guaranteed because Go maps are unordered collections.

---

### 6. **Finding vowels in a string (with indices)** (`main5` function):
```go
func main5() {
	title := "I am learning the fundamentals of Golang"

	for index, letter := range title {
		letterString := string(letter)
		if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
			fmt.Println("index:", index, "letter:", letterString)
		}
	}
}
```
- **Purpose:** Finds and prints the vowels (`a, e, i, o, u`) in a string along with their indices.
- **How it works:**
  - The **`range`** keyword is used to iterate through the string `title`.
  - For each character (converted to `letterString`), it checks if it is a vowel (a, e, i, o, or u).
  - If it is a vowel, the loop prints the index of the vowel and the vowel itself.

**Output:**
```
index: 1 letter: a
index: 3 letter: a
index: 5 letter: e
index: 8 letter: a
index: 10 letter: i
index: 14 letter: e
index: 19 letter: a
index: 21 letter: e
index: 24 letter: u
index: 30 letter: o
index: 33 letter: a
```

---

### 7. **Finding vowels in a string (using `switch-case`)** (`main6` function):
```go
func main6() {
	title := "I am learning the fundamentals of Golang"

	for index, letter := range title {
		letterString := string(letter)
		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("index:", index, " letter:", letterString)
		}
	}
}
```
- **Purpose:** Finds and prints the vowels (`a, e, i, o, u`) in a string along with their indices using a `switch-case` statement.
- **How it works:**
  - The **`range`** keyword is used to iterate through the string `title`, just like in `main5`.
  - Instead of using an `if` statement, a `switch` statement is used to check if the character is a vowel. 
  - If the character is a vowel, the loop prints the index and the vowel.

**Output:**
```
index: 1  letter: a
index: 3  letter: a
index: 5  letter: e
index: 8  letter: a
index: 10 letter: i
index: 14 letter: e
index: 19 letter: a
index: 21 letter: e
index: 24 letter: u
index: 30 letter: o
index: 33 letter: a
```

---

### Summary:
- These functions showcase how to iterate over different data structures in Go (slices, arrays, maps, and strings) using both **traditional `for` loops** and **`range` loops**.
- The **`range`** keyword simplifies the iteration process, especially when working with collections like slices, arrays, and maps, as it directly provides both the index and the value of each element.
- The program also demonstrates how to check conditions like vowels in a string, using both **`if` statements** and **`switch-case`** blocks.