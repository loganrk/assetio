package aes

import (
	"assetio/internal/port"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type aesCipher struct {
	cryptoKey string
}

// New initializes a new aesCipher instance with the provided cryptographic key.
// This is the constructor function that returns a Cipher interface.
func New(cryptoKey string) port.Cipher {
	return &aesCipher{
		cryptoKey: cryptoKey, // Assign the cryptoKey to the aesCipher instance
	}
}

// GetKey returns the cryptographic key used in the aesCipher instance.
// This key is used for encryption and decryption operations.
func (a *aesCipher) GetKey() string {
	return a.cryptoKey // Return the stored cryptographic key
}

// Encrypt encrypts the given plaintext string using AES encryption with CFB mode.
// It first converts the text to base64 encoding, then encrypts it and returns the base64-encoded ciphertext.
func (a *aesCipher) Encrypt(text string) (string, error) {
	// Create a new AES cipher block from the cryptoKey
	block, err := aes.NewCipher([]byte(a.cryptoKey))
	if err != nil {
		return "", err // Return an error if the AES cipher block creation fails
	}

	// Encode the input text as base64 to be used in the encryption
	b := base64.StdEncoding.EncodeToString([]byte(text))
	// Create a buffer for the ciphertext with size block size + length of the base64-encoded text
	ciphertext := make([]byte, aes.BlockSize+len(b))
	// The initialization vector (IV) is the first AES.BlockSize bytes of the ciphertext
	iv := ciphertext[:aes.BlockSize]
	// Fill the IV with random bytes
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err // Return an error if the random IV generation fails
	}

	// Create the AES encryption stream in CFB mode
	stream := cipher.NewCFBEncrypter(block, iv)
	// Encrypt the base64-encoded text and store it in the ciphertext buffer (after the IV)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	// Return the final ciphertext as base64 encoding
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the provided ciphertext string back into the original plaintext.
// It first decodes the base64-encoded ciphertext, then decrypts it and returns the original plaintext.
func (a *aesCipher) Decrypt(cryptoText string) (string, error) {
	// Decode the base64-encoded ciphertext
	ciphertext, _ := base64.StdEncoding.DecodeString(cryptoText)

	// Create a new AES cipher block from the cryptoKey
	block, err := aes.NewCipher([]byte(a.cryptoKey))
	if err != nil {
		return "", err // Return an error if the AES cipher block creation fails
	}

	// Check if the ciphertext is large enough (it must be at least the block size)
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short") // Return an error if the ciphertext is too short
	}
	// The IV is stored in the first AES.BlockSize bytes of the ciphertext
	iv := ciphertext[:aes.BlockSize]
	// The actual ciphertext follows the IV
	ciphertext = ciphertext[aes.BlockSize:]

	// Create the AES decryption stream in CFB mode
	stream := cipher.NewCFBDecrypter(block, iv)
	// Decrypt the ciphertext using XOR operation
	stream.XORKeyStream(ciphertext, ciphertext)

	// Decode the decrypted text from base64 back to original data
	data, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return "", err // Return an error if base64 decoding fails
	}

	// Return the decrypted plaintext as a string
	return string(data), nil
}
