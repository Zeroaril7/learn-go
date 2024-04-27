package ngide

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func Encrypt(text, key string) error {
	eText := []byte(text)
	eKey := []byte(key)

	c, err := aes.NewCipher(eKey)

	if err != nil {
		fmt.Println(err)
		return err
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
		return err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
		return err
	}

	encryptText := gcm.Seal(nonce, nonce, eText, nil)

	fmt.Println(encryptText)

	err = os.WriteFile("encrypt.txt", encryptText, 0777)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func Decrypt(path, key string) error {
	eKey := []byte(key)
	cipherText, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return err
	}

	c, err := aes.NewCipher(eKey)

	if err != nil {
		fmt.Println(err)
		return err
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
		return err
	}

	nonceSize := gcm.NonceSize()

	if len(cipherText) < nonceSize {
		fmt.Println(err)
	}

	nonce, cipcipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipcipherText, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(plaintext))

	return err
}
