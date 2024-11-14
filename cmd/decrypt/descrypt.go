package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	cipherAes "assetio/internal/adapters/cipher/aes"
)

// main function is the entry point of the application, handling user input and decryption tasks.
func main() {
	// Create a new reader to read from standard input (keyboard).
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter the cryptographic key for decryption.
	fmt.Print("Enter the crypto key: ")
	// Read the crypto key from the user input, stopping at a newline character.
	cryptoKey, _ := reader.ReadString('\n')
	// Trim any leading or trailing spaces (including newline characters) from the input.
	cryptoKey = strings.TrimSpace(cryptoKey)

	// Initialize a new cipher instance for AES decryption using the provided crypto key.
	cipherIns := cipherAes.New(cryptoKey)

	// Enter an infinite loop to repeatedly accept text input and decrypt it.
	for {
		// Prompt the user to enter the text to decrypt.
		fmt.Print("Enter text to decrypt: ")
		// Read the ciphertext text to decrypt from the user input, stopping at a newline character.
		ciphertext, _ := reader.ReadString('\n')
		// Trim any leading or trailing spaces (including newline characters) from the input.
		ciphertext = strings.TrimSpace(ciphertext)

		// Attempt to decrypt the entered ciphertext using the cipher instance.
		decrypted, err := cipherIns.Decrypt(ciphertext)
		// If an error occurs during decryption, log the error and terminate the program.
		if err != nil {
			log.Fatalf("Error decrypting: %v", err)
			return
		}
		// Print the decrypted text to the console.
		fmt.Printf("Decrypted: %s\n", decrypted)
	}
}
