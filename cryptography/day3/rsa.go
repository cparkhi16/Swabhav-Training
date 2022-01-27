package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
)

func main() {
	privatekey, _ := rsa.GenerateKey(rand.Reader, 2048)
	fmt.Println("private key ", privatekey)
	publickey := privatekey.PublicKey
	msg := "chinmay parkhi"
	cipher, _ := rsa.EncryptOAEP(sha512.New(), rand.Reader, &publickey, []byte(msg), nil)
	fmt.Println("encrypted text ", string(cipher))
	plainText, _ := privatekey.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA512})
	fmt.Println("Plain text ", string(plainText))
}
