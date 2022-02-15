package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

type User struct {
	name      string
	BLPLevel  int //confidentiality level
	BIBALevel int //integrity level
	password  string
}

type File struct {
	name      string
	BLPLevel  int //confidentiality level
	BIBALevel int //integrity level
}

//higher level value means higher level of confidentiality
var file1 = File{name: "file1.txt", BLPLevel: 1, BIBALevel: 3}
var file2 = File{name: "file2.txt", BLPLevel: 2, BIBALevel: 2}
var file3 = File{name: "file3.txt", BLPLevel: 3, BIBALevel: 1}
var metadata = []File{file1, file2, file3}
var users = []User{User{name: "yog", BLPLevel: 1, BIBALevel: 3, password: "8c02b109849543510c1f1010d6403013"},
	User{name: "shan", BLPLevel: 3, BIBALevel: 1, password: "7be1e006cfb94fba5c8ce8da6f317ed7"}}
var salt string = "yash"

func createHashUsingMD5(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key + salt))
	//fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHashUsingMD5(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHashUsingMD5(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce := data[:nonceSize]
	cipherText := data[nonceSize:]
	plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
	return plainText
}

func getUserFromName(username string) User {
	for _, user := range users {
		if username == user.name {
			return user
		}
	}
	return User{}
}

func validateLoginCredentials(user User, password string) bool {
	emptyuser := User{}
	if user != emptyuser {
		if createHashUsingMD5(password+user.name+strconv.Itoa(user.BLPLevel)+strconv.Itoa(user.BIBALevel)) == user.password {
			return true
		} else {
			return false
		}
	}
	return false
}

func scanString(s *bufio.Scanner) (string, error) {
	if s.Scan() {
		return s.Text(), nil
	}
	err := s.Err()
	if err == nil {
		err = io.EOF
	}
	return "", err
}

func scanInt(s *bufio.Scanner) (int, error) {
	if s.Scan() {
		text := strings.TrimSpace(s.Text())
		i, err := strconv.Atoi(text)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	err := s.Err()
	if err == nil {
		err = io.EOF
	}
	return 0, err
}

func main() {
	//shan:=User{name:"shan",level:1}
	s := bufio.NewScanner(os.Stdin)
	var username string
	var operation string
	var password string
	fmt.Println("Enter your name-")
	fmt.Scanln(&username)
	fmt.Println("Enter your password-")
	fmt.Scanln(&password)
	currentUser := getUserFromName(username)
	if validateLoginCredentials(currentUser, password) {
		fmt.Println("--Logged in--")
		fmt.Println("Enter the operation(w,r)-")
		fmt.Scanln(&operation)
		accessibleFiles := getBLPAndBIBAAccessibleFiles(currentUser, operation)
		fmt.Println("list of files Accessible to you-")
		for i, file := range accessibleFiles {
			fmt.Println(i, " ", file.name)
		}
		fmt.Println("Enter the file you want to access(0,1,2,..)-")
		fileNumber, _ := scanInt(s)
		if operation == "w" {
			var data string
			fmt.Println("Enter data to write-")
			data, _ = scanString(s)
			err := WriteToFile(accessibleFiles[fileNumber].name, fs.FileMode(os.O_WRONLY), data) //change append to write
			if err != nil {
				fmt.Println("error in writing to file-", err)
			}
		} else {
			data, err := ReadFromFile(accessibleFiles[fileNumber].name)
			if err != nil {
				fmt.Println("error in writing to file-", err)
			}
			fmt.Println("file data-", data)
		}

	} else {
		fmt.Println("Invalid login credentials")
	}

}

func getBLPAndBIBAAccessibleFiles(user User, operation string) []File {
	accessibleFiles := []File{}
	for _, file := range metadata {
		blpCheck := !((file.BLPLevel > user.BLPLevel && operation == "r") || (file.BLPLevel < user.BLPLevel && operation == "w"))
		bibaCheck := !((file.BIBALevel > user.BIBALevel && operation == "w") || (file.BIBALevel < user.BIBALevel && operation == "r"))
		if blpCheck && bibaCheck {
			accessibleFiles = append(accessibleFiles, file)
		}
	}
	return accessibleFiles
}

func WriteToFile(filename string, fileMode os.FileMode, data string) error {
	file, err := os.OpenFile("./data/"+filename, int(fileMode), 0777)
	fileContent, _ := ReadFromFile(filename)
	//fmt.Println(fileContent,"---",data,string(encrypt([]byte(fileContent+data),"hello")))
	if err != nil {
		//fmt.Println(err)
		return err
	}
	defer file.Close()
	if _, err := file.Write(encrypt([]byte(fileContent+data), "hello")); err != nil {
		return err
	}
	return nil
}

func ReadFromFile(filename string) (string, error) {
	data, err := os.ReadFile("./data/" + filename)
	//fmt.Println(data)
	if err != nil {
		return "", err
	}
	if string(data) == "" {
		return "", nil
	}
	return string(decrypt(data, "hello")), nil
}
