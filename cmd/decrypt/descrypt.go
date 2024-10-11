package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/loganrk/go-cipher"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the crypto key: ")
	cryptoKey, _ := reader.ReadString('\n')
	cryptoKey = strings.TrimSpace(cryptoKey)
	cipherIns := cipher.New(cryptoKey)

	for {
		fmt.Print("Enter text to decrypt: ")
		ciphertext, _ := reader.ReadString('\n')
		ciphertext = strings.TrimSpace(ciphertext)

		decrypted, err := cipherIns.Decrypt(ciphertext)
		if err != nil {
			log.Fatalf("Error decrypting: %v", err)
			return
		}
		fmt.Printf("Decrypted: %s\n", decrypted)
	}
}
