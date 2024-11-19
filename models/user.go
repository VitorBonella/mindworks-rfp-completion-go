package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" gorm:"unique"`
	Password []byte `json:"-"`
	ApiKey string `json:"-"`
	HasApiKey bool `gorm:"-"`
}

// Function to encrypt the API Key
func encrypt(key, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Function to decrypt the API Key
func decrypt(key, cryptoText string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(cryptoText)
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

	return string(ciphertext), nil
}

// Function to set an encrypted API Key for a User
func (u *User) SetAPIKey(apiKey string) error {
	masterKey := os.Getenv("MASTER_KEY")
	encryptedKey, err := encrypt(masterKey, apiKey)
	if err != nil {
		return err
	}
	u.ApiKey = encryptedKey
	return nil
}

// Function to get the decrypted API Key for a User
func (u *User) GetAPIKey() (string, error) {
	masterKey := os.Getenv("MASTER_KEY")
	return decrypt(masterKey, u.ApiKey)
}