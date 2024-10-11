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

func New(cryptoKey string) port.Cipher {
	return &aesCipher{
		cryptoKey: cryptoKey,
	}
}

func (a *aesCipher) GetKey() string {
	return a.cryptoKey
}

func (a *aesCipher) Encrypt(text string) (string, error) {

	block, err := aes.NewCipher([]byte(a.cryptoKey))
	if err != nil {
		return "", err
	}

	b := base64.StdEncoding.EncodeToString([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return base64.StdEncoding.EncodeToString(ciphertext), nil

}

func (a *aesCipher) Decrypt(cryptoText string) (string, error) {

	ciphertext, _ := base64.StdEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher([]byte(a.cryptoKey))
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	data, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return "", err
	}

	return string(data), nil
}
