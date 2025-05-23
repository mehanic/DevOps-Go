package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	// Use a 16-byte key for AES-128
	key := []byte("my-secret-key-12")
	plaintext := []byte("Hello, world!")

	// Create a new AES cipher block with the given key
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Prepare the ciphertext slice with extra space for the IV
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Generate a random IV
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a new CTR stream cipher and encrypt the plaintext
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// Output the ciphertext and plaintext
	fmt.Println("Ciphertext:", ciphertext)
	fmt.Println("Plaintext:", plaintext)
}


// go run encription.go                                                                                                                                          
// Ciphertext: [211 181 139 27 15 9 31 55 33 190 223 193 97 194 109 227 45 45 160 177 99 183 135 165 204 28 191 168 180]
// Plaintext: [72 101 108 108 111 44 32 119 111 114 108 100 33]
