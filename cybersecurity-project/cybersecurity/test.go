package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func DesEncryption(key, iv, plainText []byte) ([]byte, error) {

	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := PKCS5Padding(plainText, blockSize) //des uses 64 bit key and we cannot use gcm here so we use PKCS5Padding
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}

func DesDecryption(key, iv, cipherText []byte) ([]byte, error) {

	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func adjustSize(text []byte) []byte {
	//fmt.Println(len(text))
	if len(text) >= 8 {
		return text[:8]
	} else {
		return PKCS5Padding(text, 8)
	}
}

func main() {
	var originalText string
	fmt.Println("Enter Original text-")
	fmt.Scanf("%s\n", &originalText)
	//originalText := "wort"
	fmt.Println("originalText", originalText)
	mytext := []byte(originalText)

	key := []byte("wer")
	iv := []byte("asd4444444444444444444444")
	key = adjustSize(key)
	iv = adjustSize(iv)
	//fmt.Println(key, iv)

	cryptoText, _ := DesEncryption(key, iv, mytext)
	fmt.Println("en", string(cryptoText))
	decryptedText, _ := DesDecryption(key, iv, cryptoText)
	fmt.Println("dec", string(decryptedText))

}
