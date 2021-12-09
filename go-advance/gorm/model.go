package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	gorm.Model
	Name    string
	Address string `gorm:"index:addr"` // create index with name `addr` for address
	Gender  string
	IsMale  bool
}

func GetUsers(db *gorm.DB) {
	var p []Person
	db.Debug().Find(&p)
	fmt.Println("DB ALL PEOPLE ", p)
}
func main() {
	dbConn := "root:hello@tcp(127.0.0.1:3306)/newdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	//db.AutoMigrate(&Person{})
	//p := Person{Name: "Chinmay", Address: "Dom", Gender: "Male", IsMale: true}
	//db.Debug().Create(&p)

	//p2 := Person{Name: "Key", Address: "Dom", Gender: "Male", IsMale: true}
	//db.Debug().Create(&p2)

	//p3 := Person{Name: "Ram", Address: "Tha", Gender: "Male", IsMale: true}
	//db.Debug().Create(&p3)
	GetUsers(db)
	//var p4 Person
	//db.Model(&p4).Where("name = ?", "Key").Debug().Update("name", "Keyur")
	var pt Person
	db.Where("name = ?", "Keyur").Debug().Find(&pt)
	fmt.Println("Details of person ", pt)

	//db.Debug().Delete(&pt)
	db.Unscoped().Where("name = ?", "Keyur").Debug().Find(&pt)
	fmt.Println(pt)
}
