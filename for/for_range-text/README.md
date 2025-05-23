This Go program counts the frequency of each word in a given text. Let's break it down step by step to understand how it works:

### Code Breakdown

#### 1. **Function Definition: `countWord`**
```go
func countWord(text string) map[string]int {
```
- The function `countWord` takes a string (`text`) as input and returns a map where the keys are words (of type `string`) and the values are their frequencies (of type `int`).

#### 2. **Convert Text to Lowercase and Split Into Words**
```go
words := strings.Fields(strings.ToLower(text))
```
- `strings.ToLower(text)` converts the entire input text to lowercase. This ensures that the word count is case-insensitive (e.g., "The" and "the" are treated as the same word).
- `strings.Fields(text)` splits the lowercase text into a slice of words, using spaces as delimiters.

#### 3. **Create a Map to Track Word Frequencies**
```go
frq := make(map[string]int, len(words))
```
- A map `frq` is created to store the word frequency count. The map's key type is `string` (to store words), and the value type is `int` (to store the frequency of each word).
- `len(words)` is passed as the second argument to `make` to preallocate space for the map, based on the number of words in the input text.

#### 4. **Iterate Over Each Word**
```go
for _, word := range words {
```
- This loop iterates through each word in the `words` slice.

#### 5. **Remove Punctuation from Words**
```go
word = strings.Trim(word, `.-,"!?`)
```
- The function `strings.Trim` is used to remove common punctuation marks (`.-,"!?`) from the start and end of each word. This helps ensure that words like "hello," and "hello." are treated as the same word ("hello").
  
#### 6. **Update Word Frequency Count**
```go
frq[word]++
```
- For each cleaned word, its count is incremented in the `frq` map. If the word already exists in the map, its value is incremented by 1. If the word is not found in the map, it is added with an initial count of 1.

#### 7. **Return the Frequency Map**
```go
return frq
```
- After processing all the words, the function returns the map `frq`, which contains the frequency of each word in the input text.

#### 8. **Main Function**
```go
func main() {
	test_text := `As far as eye could reach he saw nothing but the stems of the great plants about him receding 
	in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the 
	solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued 
	soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. 
	Once or twice a small red creature scuttled across his path,
	but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering 
	unprovisioned and alone in a forest of unknown vegetation thousands or millions of 
	miles beyond the reach or knowledge of man.
	`
```
- This block contains the input text (`test_text`) for which the word frequency count will be calculated. This is a short passage that is used as the example input.

#### 9. **Calling the `countWord` Function**
```go
frq := countWord(test_text)
```
- The `countWord` function is called with the `test_text` string, and the result (the map containing word frequencies) is assigned to the variable `frq`.

#### 10. **Print Word Frequencies**
```go
for word, count := range frq {
    fmt.Printf("%d -- %v\n", count, word)
}
```
- This loop iterates over the `frq` map, printing each word and its corresponding frequency. The output is formatted to show the frequency of each word followed by the word itself.

### Example Output
After running the program, the output will look like this (the actual output will depend on the word count in `test_text`):

```
1 -- as
1 -- far
1 -- eye
...
```

The output shows the frequency of each word in the given text. Words that appear more frequently will have a higher count.

### Key Points:
1. **Case Insensitivity**: The program converts the entire text to lowercase to ensure the word count is case-insensitive.
2. **Punctuation Removal**: The program uses `strings.Trim` to remove common punctuation marks around words.
3. **Word Frequency Count**: It counts how many times each word appears and stores it in a map.
4. **Efficient Iteration**: The program iterates over the words and updates the frequency map using `map[word]++`.

### Edge Cases Handled:
- The program ensures that punctuation marks at the beginning or end of words are removed.
- It handles case-insensitive comparisons by converting the entire text to lowercase before splitting it into words.
