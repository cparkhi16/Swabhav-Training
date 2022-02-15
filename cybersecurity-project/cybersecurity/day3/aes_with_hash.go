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

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	//fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce := data[:nonceSize]
	cipherText := data[nonceSize:]
	plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
	return plainText
}

func receivedMsgIsOk(msg string) bool {
	msgHash := msg[:32]
	//fmt.Println(msgHash)
	msgText := msg[32:]
	msgHash2 := createHash(msgText)
	//fmt.Println(msgHash, msgHash2)
	if msgHash2 == msgHash {
		return true
	}
	return false
}

func main() {
	cipherText := encrypt([]byte("yogesh"), "hello")
	cipherTextHash := createHash(string(cipherText))
	newCipherText := cipherTextHash + string(cipherText)
	fmt.Println("message sent- ", newCipherText)
	if receivedMsgIsOk(newCipherText) {
		fmt.Println("message received- ", string(decrypt(cipherText, "hello")))
	}
}
