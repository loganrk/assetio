package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	cipherAes "assetio/internal/adapters/cipher/aes"
)

// main function is the entry point of the application, handling user input and encryption tasks.
func main() {
	// Create a new reader to read from standard input (keyboard).
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter the cryptographic key for encryption.
	fmt.Print("Enter the crypto key: ")
	// Read the crypto key from the user input, stopping at a newline character.
	cryptoKey, _ := reader.ReadString('\n')
	// Trim any leading or trailing spaces (including newline characters) from the input.
	cryptoKey = strings.TrimSpace(cryptoKey)

	// Initialize a new cipher instance for AES encryption using the provided crypto key.
	cipherIns := cipherAes.New(cryptoKey)

	// Enter an infinite loop to repeatedly accept text input and encrypt it.
	for {
		// Prompt the user to enter the text to encrypt.
		fmt.Print("Enter text to encrypt: ")
		// Read the plaintext text to encrypt from the user input, stopping at a newline character.
		plaintext, _ := reader.ReadString('\n')
		// Trim any leading or trailing spaces (including newline characters) from the input.
		plaintext = strings.TrimSpace(plaintext)

		// Attempt to encrypt the entered plaintext using the cipher instance.
		encrypted, err := cipherIns.Encrypt(plaintext)
		// If an error occurs during encryption, log the error and terminate the program.
		if err != nil {
			log.Fatalf("Error encrypting: %v", err)
			return
		}
		// Print the encrypted text to the console.
		fmt.Printf("Encrypted: %s\n", encrypted)
	}
}
