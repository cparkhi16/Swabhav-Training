package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
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
	fmt.Println(data)
	key := []byte(createHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce := data[:nonceSize]
	cipherText := data[nonceSize:]
	plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
	fmt.Println("Pt ", plainText)
	return plainText
}
func WriteFile(filename string, fileMode os.FileMode, data string) error {
	d, _ := ReadFile(filename)
	dat := string(Decrypt([]byte(d), "hellohellohelloo"))
	fmt.Println("Decrypted data ", dat)
	// file, err := os.OpenFile(filename, int(fileMode), 0777)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()
	// if _, err := file.Write([]byte(Encryption([]byte(data), "hellohellohelloo"))); err != nil {
	// 	return err
	// }
	return nil
}

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	fmt.Println("Data here ", data)
	//fmt.Println("---", aes.Decryption(data, "hello"))
	if err != nil {
		return "", err
	}
	return string(Decrypt(data, "hellohellohelloo")), nil
}
func main() {
	// file := model.NewFile("employee.csv", "2") // no write // BELL
	// AddFiles(*file)
	// fileTwo := model.NewFile("items.csv", "4") // no read //BELL
	// AddFiles(*fileTwo)
	// fmt.Println("1.Register 2.Login ")
	// var response string
	// fmt.Scanln(&response)
	// if response == "1" {
	// 	fmt.Println("Enter name password biba and bell levels")
	// 	var name string
	// 	var password string
	// 	var biba string
	// 	var bell string
	// 	fmt.Scanf("%s %s %s %s", &name, &password, &biba, &bell)
	// 	user := model.NewUser(name, password, biba, bell)
	// 	user.AddUser()
	// 	mode := "BIBA"
	// 	CheckWriteAccessForFiles(user, mode)
	// 	CheckReadAccessForFiles(user, mode)
	// } else {
	// 	//user := model.NewUser("Chinmay", "test", "3", "3")
	// 	user := model.Login()
	// 	if user == nil {
	// 		fmt.Println("Invalid username or password ")
	// 	} else {
	// 		mode := "BELL"
	// 		CheckWriteAccessForFiles(user, mode)
	// 		CheckReadAccessForFiles(user, mode)
	// 	}
	// }
	// employeeCSV()
	//itemsCSV()
	// user := model.NewUser("Chinmay", "test", "3", "3")
	// CheckWriteAccessForFiles(user)
	// CheckReadAccessForFiles(user)
	var data string
	fmt.Println("Enter data")
	fmt.Scan(&data)
	err := WriteFile("test.txt", fs.FileMode(os.O_APPEND), data)
	if err != nil {
		fmt.Println("-= Error while writing to file -=", err)
	}
	d, err := ReadFile("test.txt")
	if err != nil {
		fmt.Println("-= Error reading file -=", err)
	}
	fmt.Println("Contents of  file ", d)
}
