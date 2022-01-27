package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func createHashUsingMD5(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	//fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encryptAndHash(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHashUsingMD5(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	fmt.Println("Len of byte slice of cipherText ", len(cipherText), cipherText)
	cipherTextHash := createHashUsingMD5(string(cipherText))
	fmt.Println("Hash of cipher text ", cipherTextHash, len(cipherTextHash))
	dataToBeSent := cipherTextHash + string(cipherText)
	// fmt.Println("only hash value ", dataToBeSent[:32])
	// fmt.Println("Actual cipher text ", dataToBeSent[32:])
	return []byte(dataToBeSent)
}

func checkIntegrity(data []byte, passphrase string) []byte {
	key := []byte(createHashUsingMD5(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	cipherTextHash := data[:32]
	data = data[32:]

	nonce := data[:nonceSize]
	//dataToBeChecked := data[nonceSize:]
	//fmt.Println("Data ", dataToBeChecked)
	cipherText := data[nonceSize:]
	fmt.Println("CipherText ", cipherText)

	fmt.Println("Hash value of ciphertext in msg ", string(cipherTextHash))
	fmt.Println("Creating ", createHashUsingMD5(string(data)))
	if string(cipherTextHash) == createHashUsingMD5(string(data)) {
		fmt.Println("hash confirm")
	}
	plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
	fmt.Println("Plain text decrypted ", string(plainText))
	return plainText
}

func main() {
	cipherText := encryptAndHash([]byte("Chinmay"), "hello")
	fmt.Println(cipherText)
	fmt.Println(string(checkIntegrity(cipherText, "hello")))
}
