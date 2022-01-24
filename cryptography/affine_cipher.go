package main

import (
	ac "algo/cipherAlgo"
	"fmt"
	"strings"
)

func main() {
	pt := "chinmay parkhi"
	pt = strings.ReplaceAll(pt, " ", "")
	mulKey := 7
	addKey := 1
	ct := ac.MultiplicativeEncryption(pt, mulKey)
	fmt.Println("after mul ", ct)
	result := strings.Map(func(r rune) rune {
		return ac.Caesar(r, addKey)
	}, ct)
	fmt.Println("Cipher text ", result)
	decryptResult := strings.Map(func(r rune) rune {
		return ac.Decrypt(r, addKey)
	}, result)
	fmt.Println("Decrypted Result after subtract ", decryptResult)
	decryptedTextResult := ac.MultiplicativeDecryption(decryptResult, mulKey, ac.GetInverse(mulKey))
	fmt.Println("Final decypted text ", decryptedTextResult)
}
