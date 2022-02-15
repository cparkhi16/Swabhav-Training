package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	dataSourceName := "root:Panda@19@tcp(localhost:3306)/goadv?parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	fmt.Println(db)

}
