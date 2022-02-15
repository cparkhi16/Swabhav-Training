package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Employee struct {
	Name    string
	Id      int
	Address string
	Gender  string
}

func main() {
	dataSourceName := "root:****@tcp(localhost:3306)/goadv?parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	fmt.Println(db)

	//db.Exec("USE goadv")
	db.AutoMigrate(&Employee{})
	user := Employee{Name: "Jinzhu", ID: 2, Address: "Mars"}

	result := db.Create(&user) // pass pointer of data to Create

	fmt.Println(result.RowsAffected) // returns inserted records count

}
