This Go program demonstrates the use of **if-else statements** and **command-line arguments** (`os.Args`). Let's go step by step and explain each part.

---

## **1. `main()` Function**
```go
func main() {
	age := 10

	if age > 60 {
		fmt.Println("Getting older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Young blood")
	} else {
		fmt.Println("Booting up")
	}
}
```
- It initializes `age = 10`.
- It checks conditions in sequence:
  - `age > 60`: "Getting older"
  - `age > 30`: "Getting wiser"
  - `age > 20`: "Adulthood"
  - `age > 10`: "Young blood" ✅ (this condition is `true`, so it prints this and exits).
  - If none of the above were true, it would print `"Booting up"`.

**Expected Output**:
```
Young blood
```

---

## **2. `main1()` - Checking Multiple Conditions**
```go
func main1() {
	isSphere, radius := true, 200

	if isSphere && radius >= 200 {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}
```
- `isSphere = true`, `radius = 200`
- Since `isSphere && radius >= 200` is `true`, it prints:
```
It's a big sphere.
```

---

## **3. `main2()` - Checking Command-Line Arguments**
```go
func main2() {
	var (
		args = os.Args
		l    = len(args) - 1
	)

	if l == 0 {
		fmt.Println("Give me args")
	} else if l == 1 {
		fmt.Printf("There is one: %q\n", args[1])
	} else if l == 2 {
		fmt.Printf(
			`There are two: "%s %s"`+"\n",
			args[1], args[2],
		)
	} else {
		fmt.Printf("There are %d arguments\n", l)
	}
}
```
- `os.Args` captures command-line arguments.
- If no arguments are provided, it prints `"Give me args"`.
- If exactly **one** argument is given, it prints `"There is one: <arg>"`.
- If exactly **two** arguments are given, it prints `"There are two: "<arg1> <arg2>""`.
- If **more than two** arguments are given, it prints `"There are X arguments"`.

**Example Run & Outputs**:
```
$ go run main.go
Give me args

$ go run main.go hello
There is one: "hello"

$ go run main.go hello world
There are two: "hello world"

$ go run main.go a b c
There are 3 arguments
```

---

## **4. `main3()` & `main4()` - Checking Vowels and Consonants**
Both functions check if a given character is a **vowel, consonant, or special case (`w`/`y`)**.

### **main3() - Using Multiple `if` Conditions**
```go
func main3() {
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	s := args[1]
	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}
}
```
- Takes a single letter as input.
- Checks if it’s a vowel (`a, e, i, o, u`).
- Checks if it's sometimes a vowel (`w, y`).
- Otherwise, it's a consonant.

**Example Run**:
```
$ go run main.go a
"a" is a vowel.

$ go run main.go y
"y" is sometimes a vowel, sometimes not.

$ go run main.go k
"k" is a consonant.
```

### **main4() - Using `strings.IndexAny()`**
```go
func main4() {
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	s := args[1]
	if strings.IndexAny(s, "aeiou") != -1 {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}
}
```
- Uses `strings.IndexAny(s, "aeiou")` to check if `s` is a vowel more efficiently.

---

## **5. `main5()` & `main6()` - Age Classification**
These functions categorize users based on their age.

### **main5() - Using Nested `if`**
```go
func main5() {
	if len(os.Args) != 2 {
		fmt.Println("Requires age")
		return
	}

	age, err := strconv.Atoi(os.Args[1])

	if err != nil || age < 0 {
		fmt.Printf(`Wrong age: %q`+"\n", os.Args[1])
		return
	} else if age > 17 {
		fmt.Println("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	}
}
```
- Takes an age as input.
- If the argument is missing or invalid, it prints an error.
- If age > 17 → `"R-Rated"`
- If age is **13-17** → `"PG-13"`
- If age < 13 → `"PG-Rated"`

**Example Run**:
```
$ go run main.go 18
R-Rated

$ go run main.go 15
PG-13

$ go run main.go 10
PG-Rated

$ go run main.go abc
Wrong age: "abc"
```

### **main6() - Using Inline `if` (Not Recommended)**
```go
func main6() {
	if len(os.Args) != 2 {
		fmt.Println("Requires age")
		return
	} else if age, err := strconv.Atoi(os.Args[1]); err != nil || age < 0 {
		fmt.Printf(`Wrong age: %q`+"\n", os.Args[1])
		return
	} else if age > 17 {
		fmt.Println("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	}
}
```
- The **entire logic is crammed into one if-else block**.
- Harder to read, but works the same as `main5()`.

---

## **Summary of Rules**
1. **Basic `if-else`**: Checks age ranges and prints messages.
2. **Boolean Conditions**: `if isSphere && radius >= 200` checks multiple conditions.
3. **Command-line Arguments (`os.Args`)**:
   - `len(os.Args)` determines how many arguments were passed.
   - `strconv.Atoi()` converts string to integer.
   - `strings.IndexAny()` checks for characters in a string.
4. **Scope in `if` Statements**:
   - Inline variable declarations (`if age, err := strconv.Atoi(); ...`) limit variables to the `if` block.

---

Let me know if you have any questions! 🚀