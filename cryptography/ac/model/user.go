package model

import (
	"app/hash"
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type User struct {
	Name      string
	LevelBIBA string
	LevelBELL string
	Password  string
}

func NewUser(name, Password, LevelBIBA, LevelBELL string) *User {
	//fmt.Println("name ", name)
	Password = hash.CreateHashForPassword(Password)
	return &User{Name: name, Password: Password, LevelBIBA: LevelBIBA, LevelBELL: LevelBELL}
}

func (user *User) AddUser() {
	csvfile, err := os.OpenFile("user.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	defer writer.Flush()
	records := [][]string{{user.Name, user.Password, string(user.LevelBELL), string(user.LevelBIBA)}}
	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}
func Login() *User {
	fmt.Println("Enter your username and password")
	var name string
	var password string
	fmt.Scanf("%s %s", &name, &password)
	file, err := os.Open("user.csv")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	// for _, eachline := range txtlines {
	// 	fmt.Println(eachline)
	// }
	for _, val := range txtlines {
		//fmt.Println(val[1])
		tmp := strings.Split(val, ",")
		currUserName := tmp[0]
		currUserPassword := tmp[1]
		if currUserName == name {
			if hash.ComparePasswords(currUserPassword, password) {
				fmt.Println("Login successfull ")
				return &User{Name: currUserName, Password: currUserPassword, LevelBIBA: tmp[2], LevelBELL: tmp[3]}
			}
		}
	}
	return nil
}
