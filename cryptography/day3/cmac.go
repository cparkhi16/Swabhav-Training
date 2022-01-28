package main

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/aead/cmac"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	//fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	return hex.EncodeToString(hasher.Sum(nil))
}
func main() {
	data := "chinmay parkhi"
	block, _ := aes.NewCipher([]byte(createHash("hello")))
	//c, _ := cmac.New(block)
	actual, _ := cmac.Sum([]byte(data), block, block.BlockSize())
	fmt.Println(actual)
	sha := string(actual)
	isValid := cmac.Verify(actual, []byte("chinmay parkhi"), block, block.BlockSize())
	fmt.Println("Is valid ", isValid)
	fmt.Println("cmac", sha)
}
