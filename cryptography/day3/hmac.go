package main

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func encryptUsingRSA(data string) (string, *rsa.PrivateKey) {
	privatekey, _ := rsa.GenerateKey(rand.Reader, 2048)
	//fmt.Println("private key ", privatekey)
	publickey := privatekey.PublicKey
	cipher, _ := rsa.EncryptOAEP(sha512.New(), rand.Reader, &publickey, []byte(data), nil)
	fmt.Println("encrypted text ", string(cipher))
	return string(cipher), privatekey
}
func HMAC(secret string, data string) string {

	fmt.Printf("Secret: %s Data: %s\n", secret, data)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Println("Result: "+sha, len(sha))
	return sha
}
func DecryptUsingRSA(privatekey *rsa.PrivateKey, data string) {
	hmac := data[:64]
	fmt.Println("HMAC in decrypt ", hmac)
	generatedMAC := HMAC("hello", data[64:])
	if generatedMAC == hmac {
		fmt.Println("Authenticated ")
		cipher := []byte(data[64:])
		plainText, _ := privatekey.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA512})
		fmt.Println("Plain text ", string(plainText))
	} else {
		fmt.Println("Not authenticated")
	}

}
func main() {
	mydata := "chinmay parkhi"
	cipher, privatekey := encryptUsingRSA(mydata)
	mac := HMAC("hello", cipher)
	finalData := mac + cipher
	//fmt.Println("Appending HMAC+ CIPHER ", finalData[:64])
	DecryptUsingRSA(privatekey, finalData)
}
