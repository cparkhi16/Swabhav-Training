package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	//fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	log.Print("File encryption example")

	plaintext, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// The key should be 16 bytes (AES-128), 24 bytes (AES-192) or
	// 32 bytes (AES-256)
	// key, err := ioutil.ReadFile("key")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	block, _ := aes.NewCipher([]byte(createHash("hello")))
	//gcm, _ := cipher.NewGCM(block)
	// block, err := aes.NewCipher()
	// if err != nil {
	// 	log.Panic(err)
	// }

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}

	// Never use more than 2^32 random nonces with a given key
	// because of the risk of repeat.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	// Save back to file
	err = ioutil.WriteFile("ciphertext.bin", ciphertext, 0777)
	if err != nil {
		log.Panic(err)
	}
}
