This Go program works with a list of champions (from a "TFT" game, likely *Teamfight Tactics*) represented by a struct `champion`. The program does several operations such as loading champion data from a JSON file, checking if champions belong to specific classes or origins, removing specific champions, and modifying slices of champions based on different criteria.

Let's go through the code, explain the various rules and operations, and how each part works.

### **1. Champion Struct Definition**

```go
type champion struct {
	Name    string   `json:"name"`
	Classes []string `json:"classes"`
	Origins []string `json:"origins"`
	Cost    int      `json:"cost"`
}
```

- This defines a `champion` struct, which contains:
  - `Name`: The name of the champion (string).
  - `Classes`: A list of classes the champion belongs to (slice of strings).
  - `Origins`: A list of origins the champion belongs to (slice of strings).
  - `Cost`: The cost of the champion (integer).

The struct uses **JSON tags** to specify the field names used when encoding or decoding from JSON (e.g., `"name"` for `Name`).

### **2. Methods for the `champion` Struct**

The following methods are defined on the `champion` struct to help check whether a champion has certain classes or origins.

#### **`hasClass` Method**

```go
func (c champion) hasClass(class string) bool {
	for _, champClass := range c.Classes {
		if champClass == class {
			return true
		}
	}
	return false
}
```

- **Purpose**: Checks if a champion belongs to a specific class.
- **Operation**: It loops through the `Classes` slice of the `champion` struct and checks if the given class exists in the list. If found, it returns `true`, otherwise `false`.

#### **`hasAllClasses` Method**

```go
func (c champion) hasAllClasses(classes ...string) bool {
	classCount := len(classes)
	foundCount := 0
	for _, class := range classes {
		if found := c.hasClass(class); found {
			foundCount++
		}
	}
	return foundCount == classCount
}
```

- **Purpose**: Checks if the champion belongs to **all** of the given classes.
- **Operation**: The method takes a variadic parameter `classes`, which allows passing any number of classes. It checks each class using the `hasClass` method, and counts how many of the classes are found. If all the classes are found, it returns `true`.

#### **`hasSomeClasses` Method**

```go
func (c champion) hasSomeClasses(classes ...string) bool {
	for _, class := range classes {
		if found := c.hasClass(class); found {
			return true
		}
	}
	return false
}
```

- **Purpose**: Checks if the champion belongs to **any** of the given classes.
- **Operation**: Similar to `hasAllClasses`, but it returns `true` as soon as it finds any class from the given list. If none of the classes are found, it returns `false`.

#### **`hasOrigin` Method**

```go
func (c champion) hasOrigin(origin string) bool {
	for _, champOrigin := range c.Origins {
		if champOrigin == origin {
			return true
		}
	}
	return false
}
```

- **Purpose**: Checks if a champion belongs to a specific origin.
- **Operation**: Similar to `hasClass`, but this method checks for an origin in the `Origins` slice instead of the `Classes` slice.

#### **`hasAllOrigins` Method**

```go
func (c champion) hasAllOrigins(origins ...string) bool {
	originCount := len(origins)
	foundCount := 0
	for _, origin := range origins {
		if found := c.hasOrigin(origin); found {
			foundCount++
		}
	}
	return foundCount == originCount
}
```

- **Purpose**: Checks if the champion belongs to **all** of the given origins.
- **Operation**: Similar to `hasAllClasses`, but for origins. It returns `true` if the champion belongs to all the specified origins.

#### **`hasSomeOrigins` Method**

```go
func (c champion) hasSomeOrigins(origins ...string) bool {
	for _, origin := range origins {
		if found := c.hasOrigin(origin); found {
			return true
		}
	}
	return false
}
```

- **Purpose**: Checks if the champion belongs to **any** of the given origins.
- **Operation**: This method works similarly to `hasSomeClasses` but works with origins.

### **3. Loading Champions from a JSON File**

```go
func loadChampions() ([]champion, error) {
	file, err := os.Open("tft_champions.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var champions []champion
	for {
		if err := json.NewDecoder(file).Decode(&champions); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	return champions, nil
}
```

- **Purpose**: Loads a list of champions from a file (`tft_champions.json`) into a slice of `champion` structs.
- **Operation**: 
  - It opens the JSON file using `os.Open`.
  - It uses `json.NewDecoder(file).Decode(&champions)` to decode the JSON data into the `champions` slice.
  - It returns the loaded slice and any error that occurs.

### **4. Main Function**

```go
func main() {
	all, err := loadChampions()
	if err != nil {
		log.Fatalf("An error occurred loading/parsing champions, err=%v", err)
	}

	fmt.Printf("There are %d/%d champions (all): %v\n\n", len(all), cap(all), all)

	first := all[0]
	fmt.Printf("The first champion is %v\n\n", first)

	some := all[0:20]
	some = all[:20]
	some = all[10:]
	some = all[:]
	fmt.Printf("The first %d champions are %v\n\n", len(some), some)

	// Remove an element from the slice by index
	most := append(all[0:10], all[11:]...)
	fmt.Printf(
		"There are %d/%d champions (most): %v\n\n",
		len(most),
		cap(most),
		most,
	)
	fmt.Printf("There are %d/%d champions (all): %v\n\n", len(all), cap(all), all)

	// Remove elements by altering the slice in place
	index := 0
	for _, champ := range all {
		if champ.hasClass("Sorcerer") {
			continue
		}

		all[index] = champ
		index++
	}

	all = all[:index]

	fmt.Printf(
		"There are %d champions after removing Sorcerers, %v\n\n",
		len(all),
		all,
	)
}
```

#### **Main Operations**:
1. **Loading Champions**:
   - The program loads champions using the `loadChampions` function.
   
2. **Printing All Champions**:
   - It prints the number of champions (`len(all)`) and their details.
   
3. **Slicing the Array**:
   - It demonstrates different slice operations:
     - `all[0:20]`: The first 20 champions.
     - `all[:20]`: Same as `all[0:20]`.
     - `all[10:]`: Champions starting from the 11th one.
     - `all[:]`: All champions.
   
4. **Removing an Element from the Slice**:
   - It removes the 11th element from the list by using `append` to concatenate two slices: `all[0:10]` (first 10 champions) and `all[11:]` (champions after the 11th one).
   
5. **Filtering Sorcerers**:
   - It iterates over all champions, removing those that have the "Sorcerer" class using `hasClass("Sorcerer")`.
   - After filtering, it prints the updated list of champions.

---

### **Summary**:
- This Go program demonstrates how to work with slices, structs, and methods to manage and manipulate data in Go.
- It defines various methods to check for specific conditions (e.g., belonging to specific classes or origins) and manipulates slices to remove elements and filter champions based on those conditions.