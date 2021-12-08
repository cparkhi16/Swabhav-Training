package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID      int
	Name    string
	Address string
	Gender  string
}

func main() {
	dbConn := "root:hello@tcp(127.0.0.1:3306)/newdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	//Read
	var users []User
	db.Find(&users)
	fmt.Println(users)
	var user1 User
	db.First(&user1, 2)
	fmt.Println(user1)
	var user2 User
	db.First(&user2)
	fmt.Println(user2)
	var user3 User
	db.Last(&user3)
	fmt.Println(user3)

	//Save It creates new record with new ID 100
	/*var userNew User
	db.First(&userNew)
	userNew.ID = 100
	db.Save(&userNew)*/

	//Save without creating a new record
	update := User{Name: "Chinmay", ID: 1, Address: "BCD"}
	db.Debug().Save(&update)

	//Delete //IF RECORD NOT PRESENT THEN IT GIVES ERROR
	/*var ud User
	IDToBeDeleted := 100
	e := db.Where("id= ?", IDToBeDeleted).Delete(&ud)
	if e != nil {
		log.Fatal("Error deleting unknown db entry")
	}*/

	//Update
	//var uUser User
	db.Table("users").Where("id = ?", 1).Debug().Update("test", "hello")
	fmt.Println("End")
}
