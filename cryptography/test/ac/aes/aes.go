package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	//fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encryption(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}
func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce := data[:nonceSize]
	cipherText := data[nonceSize:]
	plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
	return plainText
}

// func Decryption(data []byte, passphrase string) []byte {
// 	key := []byte(createHash(passphrase))
// 	block, _ := aes.NewCipher(key)
// 	gcm, _ := cipher.NewGCM(block)
// 	nonceSize := gcm.NonceSize()
// 	nonce := data[:nonceSize]
// 	cipherText := data[nonceSize:]
// 	//fmt.Println("Cipher Text ", cipherText)
// 	plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
// 	//fmt.Println("PT ", plainText)
// 	return plainText
// }
// func EncryptFile(filename string, data []byte, passphrase string) {
// 	f, _ := os.Create(filename)
// 	defer f.Close()
// 	f.Write(Encryption(data, passphrase))
// }
// func DecryptFile(filename string, passphrase string) []byte {
// 	data, _ := ioutil.ReadFile(filename)
// 	return Decryption(data, passphrase)
// }
