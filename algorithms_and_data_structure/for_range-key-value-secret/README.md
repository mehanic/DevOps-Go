This Go program implements a **Caesar Cipher decryption** for a given message. Let me break down the code and explain how it works:

### Code Breakdown

```go
package main

import "fmt"

func main() {
	msg := "L fdph, L vdz, L frqtxhuhg"  // Encrypted message

	for _, k := range msg {
		if k >= 'a' && k <= 'z' {  // Check if the character is a lowercase letter
			k -= 3  // Shift character by 3 positions back (decrypting)
			if k < 'a' {  // If the character goes below 'a', wrap around to the end of the alphabet
				k += 26
			}
		} else if k >= 'A' && k <= 'Z' {  // Check if the character is an uppercase letter
			k -= 3  // Shift character by 3 positions back (decrypting)
			if k < 'A' {  // If the character goes below 'A', wrap around to the end of the alphabet
				k += 26
			}
		}

		// Print the decrypted character
		fmt.Printf("%c", k)
	}
}
```

### Explanation

#### 1. **Message Declaration**
```go
msg := "L fdph, L vdz, L frqtxhuhg"
```
- The variable `msg` holds the encrypted message. This message is a Caesar ciphered text, where each letter has been shifted by 3 positions forward in the alphabet.

#### 2. **Looping through the Message**
```go
for _, k := range msg {
```
- The `for _, k := range msg` loop iterates over each character (`k`) in the message `msg`. The variable `k` holds the current character being processed.

#### 3. **Checking for Lowercase Letters**
```go
if k >= 'a' && k <= 'z' {
    k -= 3
    if k < 'a' {
        k += 26
    }
}
```
- This block checks if the character `k` is a lowercase letter (between 'a' and 'z').
- The character `k` is then shifted **backward by 3 positions** (`k -= 3`) to decrypt the message (since the encryption was done by shifting characters forward by 3).
- After shifting, if `k` goes below `'a'` (for example, if `k = 'a'` and we subtract 3, we get `k = '`), we wrap it around to the end of the alphabet by adding 26 (`k += 26`), effectively wrapping back around to 'z'.

#### 4. **Checking for Uppercase Letters**
```go
} else if k >= 'A' && k <= 'Z' {
    k -= 3
    if k < 'A' {
        k += 26
    }
}
```
- This block is similar to the previous one but for **uppercase letters** (between 'A' and 'Z').
- The same decryption process happens here, shifting by 3 positions backward and handling the wrap-around if necessary.

#### 5. **Printing the Decrypted Character**
```go
fmt.Printf("%c", k)
```
- After each character is decrypted, it's printed using `fmt.Printf("%c", k)`, which prints the character `k` to the console.

### Result
The encrypted message is `"L fdph, L vdz, L frqtxhuhg"`. After applying the Caesar Cipher decryption (shifting each letter 3 positions back), the decrypted message will be:
```
I came, I saw, I conquered
```

### Key Points:
- **Caesar Cipher**: This is a simple encryption technique where each letter is shifted by a fixed number (in this case, 3 positions).
- **Decryption**: To decrypt the message, the program shifts each letter **backward by 3 positions**.
- **Wrap Around**: If the letter goes past 'a' or 'A', it wraps around to 'z' or 'Z', respectively, to maintain the circular nature of the alphabet.
- **Character Handling**: The program handles both uppercase and lowercase letters separately.

### Example:

- If you have the encrypted message: `"L fdph"`, it is decoded to: `"I came"`.
- Similarly, `"L vdz"` is decoded to `"I saw"`.
- `"L frqtxhuhg"` becomes `"I conquered"`.

Thus, this program prints the decrypted message `"I came, I saw, I conquered"`.