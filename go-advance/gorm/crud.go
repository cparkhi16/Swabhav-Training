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
	test    string
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
	//update := User{Name: "Chinmay", ID: 1, Address: "BCD"}
	//db.Debug().Save(&update)

	//Delete //IF RECORD NOT PRESENT THEN IT GIVES ERROR
	/*var ud User
	IDToBeDeleted := 100
	e := db.Where("id= ?", IDToBeDeleted).Delete(&ud)
	if e != nil {
		log.Fatal("Error deleting unknown db entry")
	}*/

	//Update
	//var uUser User
	db.Table("users").Where("id = ?", 1).Debug().Update("test", "bye")
	//fmt.Println("End")

	// Update using Model
	var uUser User
	uUser.ID = 1
	db.Model(&uUser).Debug().Update("gender", "Male")

	//Updating a non existing entry
	/*e := db.Table("users").Where("id = ?", 109).Debug().Update("test", "hello")
	if e != nil {
		log.Fatal("err found")
	}*/

	//UPDATE WILL ALWAYS UPDATE THE ONLY GIVEN FIELDS BUT SAVE WILL ADD ZERO VALUES IF NOT SPECIFIED FOR
	// THAT OBJECT ...ALSO SAVE WILL CREATE NEW ENTRY IF ID IS NOT PRESENT BUT UPDATE WILL THROW ERROR IF ID
	//IS NOT PRESENT.

	var test User
	db.First(&test, 1) // find product with id 1
	fmt.Println(test)
	var test2 User
	db.First(&test2, "test = ?", "bye") // find user with test set to bye
	fmt.Println(test2)
	// Update - update all Users test value to trial
	//fmt.Println("Running user{} update with model")
	//db.Model(&User{}).Update("test", "trial").Debug()
	db.Model(&User{}).Debug().Update("address", "trial")
	test.test = "Gajendra"
	if result := db.Model(&test).Debug().Updates(test).Error; result != nil {
		log.Println("Unable to update data")
	}
	// WARNING when update with struct, GORM will only update those fields that with non blank value
	// For below Update, nothing will be updated as "", 0, false are blank values of their types
	db.Model(&test).Debug().Updates(User{Name: "", ID: 0, Address: ""})
	var caseNew User // IF ID IS CHANGED HERE IN UPDATES THEN NO CHANGE IS DONE SINCE IT IS A PRIMARY KEY
	r := db.Model(&caseNew).Where("id = ?", 3).Debug().Updates(User{Name: "Manager", ID: 35, Address: ""}).Error
	if r != nil {
		fmt.Println("Error")
	}
	// Create
	var newUser User // IF PRIMARY KEY IS NOT SPECIFIED IT GOT  CREATED BY DEFAULT AS 101
	newUser = User{Name: "Chinmay", Address: "Thane", Gender: "Male", test: "testing"}
	db.Debug().Create(&newUser)
	var n User
	n = User{Name: "JP", ID: 8, Address: "Mulund", Gender: "Male", test: "Hi"}
	db.Debug().Create(&n)
}
