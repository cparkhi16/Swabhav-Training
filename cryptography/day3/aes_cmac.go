package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/aead/cmac"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) string {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, data, nil)

	cmacHash, _ := cmac.New(block)
	cmacHash.Write(cipherText)
	hash := hex.EncodeToString(cmacHash.Sum(nil))
	fmt.Println("CMAC", string(hash))
	return hash + string(cipherText)
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce := data[32 : nonceSize+32]
	cipherText := data[32+nonceSize:]
	plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
	return plainText
}

func checkAuthenticity(msgReceived string, cmacGiven string, passphrase string) bool {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	cmacHash, _ := cmac.New(block)
	cmacHash.Write([]byte(msgReceived))
	hashObtained := hex.EncodeToString(cmacHash.Sum(nil))
	return hashObtained == cmacGiven
}

func main() {
	cipherText := encrypt([]byte("chinmay parkhi"), "test")
	fmt.Println("Encrypted text :", string(cipherText))
	if checkAuthenticity(cipherText[32:], cipherText[:32], "test") {
		decryptedData := decrypt([]byte(cipherText), "test")
		fmt.Println(string(decryptedData))
	} else {
		fmt.Println("Message authentication failed")
	}

}
