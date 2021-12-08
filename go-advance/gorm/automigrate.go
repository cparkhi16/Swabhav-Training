package main

import (
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

//Updates Gender in User table already created by CreateTable earlier ..If New fields are added then
//automigrate will add new fields to table but if deleted here from structs then it will not
//delete any field
func main() {
	dbConn := "root:hello@tcp(127.0.0.1:3306)/newdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	db.AutoMigrate(&User{})

}
