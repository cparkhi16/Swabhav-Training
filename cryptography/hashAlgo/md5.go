package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	fmt.Println(createHash("chinmay parkhi"))
}
