This Go program demonstrates how to work with JSON data, manipulate slices, and sort data in different ways. Here's a breakdown of the important concepts and rules used in the program:

### 1. **Champion Struct**
```go
type champion struct {
	Name    string   `json:"name"`
	Classes []string `json:"classes"`
	Origins []string `json:"origins"`
	Cost    int      `json:"cost"`
}
```
- **Structs** in Go are used to group related data. The `champion` struct represents a champion in a game, and it contains:
  - `Name`: A string representing the champion's name.
  - `Classes`: A slice of strings representing the champion's classes (e.g., roles or types).
  - `Origins`: A slice of strings representing the champion's origins (e.g., factions or regions).
  - `Cost`: An integer representing the gold cost of the champion.
  
- The struct tags (e.g., `json:"name"`) indicate the field name when the struct is serialized or deserialized to/from JSON.

### 2. **Loading JSON Data**
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
- This function loads the champions' data from a JSON file (`tft_champions.json`).
- It opens the file using `os.Open` and uses `json.NewDecoder(file).Decode(&champions)` to decode the JSON into a Go slice of `champion` structs.
- `json.NewDecoder(file)` creates a JSON decoder that reads the data from the file and decodes it into the `champions` slice.
- The loop continues until it hits the end of the file (`io.EOF`) or encounters an error during decoding.
  
  **Error Handling:**
  - If any errors occur while opening the file or decoding the JSON, they are returned as errors.

### 3. **Sorting Primitive Data Types**
```go
func sortPrimitives() {
	champions, err := loadChampions()
	if err != nil {
		log.Fatalf("An error occurred loading/parsing champions, err=%v", err)
	}

	var names []string
	for _, champ := range champions {
		names = append(names, champ.Name)
	}

	names = names[:10]

	log.Printf(
		"The first %d names are %v, sorted = %t\n\n",
		len(names),
		names,
		sort.StringsAreSorted(names),
	)

	sort.Strings(names)
	log.Printf(
		"The sorted names (ascending) are %v, sorted = %t\n\n",
		names,
		sort.StringsAreSorted(names),
	)

	var gold []int
	for _, champ := range champions {
		gold = append(gold, champ.Cost)
	}

	gold = gold[:10]

	log.Printf(
		"The first %d gold costs are %v, sorted = %t\n\n",
		len(gold),
		gold,
		sort.IntsAreSorted(gold),
	)

	sort.Ints(gold)
	log.Printf(
		"The sorted gold are %v, sorted = %t\n\n",
		gold,
		sort.IntsAreSorted(gold),
	)
}
```
- **Loading champions**: The `champions` slice is loaded using the `loadChampions` function.
  
- **Working with names**:
  - The program creates a new slice `names` to hold the names of the champions.
  - It iterates over the `champions` slice, appending each champion's name to the `names` slice.
  - The first 10 names are selected using `names[:10]` (slicing).
  - The `sort.StringsAreSorted(names)` checks if the names are already sorted.

- **Sorting names**:
  - The `sort.Strings(names)` function sorts the names in ascending order (alphabetically).
  - After sorting, it checks again whether the names are sorted using `sort.StringsAreSorted(names)`.

- **Working with gold costs**:
  - The program creates a new slice `gold` to hold the cost of the champions.
  - It appends each champion's cost (`champ.Cost`) to the `gold` slice.
  - The first 10 gold costs are selected using `gold[:10]`.
  - `sort.Ints(gold)` sorts the `gold` slice in ascending order.

- **Log statements**:
  - Throughout, `log.Printf` is used to output the state of the data at various points, both before and after sorting. It also checks if the data is sorted using `sort.StringsAreSorted(names)` and `sort.IntsAreSorted(gold)`.

### 4. **Sorting Slices with Custom Logic**
```go
func sortSlices() {
	champions, err := loadChampions()
	if err != nil {
		log.Fatalf("An error occurred loading/parsing champions, err=%v", err)
	}

	champs := champions[:10]

	log.Printf("The first %d champs are %v\n\n", len(champs), champs)

	// Sort champions in ascending alphabetical order by Name
	sort.Slice(champs, func(i, j int) bool {
		return champs[i].Name < champs[j].Name
	})
	log.Printf("The sorted champions (ascending by Name) are %v\n\n", champs)
}
```
- **Sorting with `sort.Slice`**:
  - `sort.Slice(champs, func(i, j int) bool {...})` sorts a slice based on a custom sorting logic. In this case, it sorts the slice of `champions` (specifically the first 10 champions) in **ascending order by the `Name` field**.
  - The `func(i, j int) bool { return champs[i].Name < champs[j].Name }` function compares two champions by their `Name` field and returns `true` if the champion at index `i` should come before the champion at index `j`.

### 5. **Main Function**
```go
func main() {
	sortSlices()
}
```
- The main function calls `sortSlices()`, which performs the loading and sorting operations. This function will log the results of sorting the champions by their names in ascending order.

### Summary of Key Concepts:
1. **Structs**: A way to group related data. The `champion` struct stores data for each champion in the game.
2. **JSON Parsing**: The `loadChampions` function loads data from a JSON file into Go structs using `json.NewDecoder(file).Decode`.
3. **Slices**: Slices in Go are dynamically sized, and they are used to store collections of data (e.g., names and costs of champions).
4. **Sorting**: The program demonstrates sorting slices using:
   - `sort.Strings` for sorting strings alphabetically.
   - `sort.Ints` for sorting integers.
   - `sort.Slice` for sorting slices with custom comparison logic.
5. **Error Handling**: The program checks for errors when loading JSON data and logs errors if any occur.

### Sample Output (Assuming the `tft_champions.json` is properly structured):
```
The first 10 names are [Aatrox Ahri Akali Ashe Bard Blitzcrank Braum Caitlyn Camille Cho'Gath Diana], sorted = false

The sorted names (ascending) are [Aatrox Ahri Akali Ashe Bard Blitzcrank Braum Caitlyn Camille Cho'Gath Diana], sorted = true

The first 10 gold costs are [1 2 3 4 5 6 7 8 9 10], sorted = false

The sorted gold are [1 2 3 4 5 6 7 8 9 10], sorted = true

The first 10 champs are [{Aatrox [Fighter] [Demacia] 4} {Ahri [Assassin] [Ionia] 5} ... ]

The sorted champions (ascending by Name) are [{Aatrox [Fighter] [Demacia] 4} {Ahri [Assassin] [Ionia] 5} ... ]
```

This program demonstrates handling JSON data, sorting slices, and logging in Go in an organized and practical way.