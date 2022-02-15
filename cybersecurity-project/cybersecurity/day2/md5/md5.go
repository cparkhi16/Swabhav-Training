package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func createHashUsingMD5(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	text := "yogesh"
	fmt.Println(createHashUsingMD5(text))
}
