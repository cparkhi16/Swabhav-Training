package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type User struct {
	Name    string
	Id      int
	Address string
}

func main() {
	dataSourceName := "root:*****@tcp(localhost:3306)/goadv?parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	fmt.Println(db)

	db.Exec("USE goadv")
	//GetAll
	var users []User
	db.Debug().Find(&users)
	fmt.Println(users)

	//GetFromId
	var user1 User
	ID := 2
	db.Debug().First(&user1, ID)
	fmt.Println(user1)

	//create

	user2 := User{Name: "user2", Id: 2, Address: "kan"}
	db.Debug().Create(&user2)
	fmt.Println(user2)

	user3 := User{Name: "user3", Id: 3, Address: "jahs"}
	db.Debug().Create(&user3)
	fmt.Println(user3)

	//update
	updatedUser := User{Name: "minzu", Id: 10}
	db.Debug().Save(&updatedUser)
	fmt.Println(updatedUser)

	//update2
	db.Debug().Table("users").Where("id = ?", 234).Update("Name", "hello")
	//fmt.Println(updatedUser1)

	//delete
	IDdelete := 2
	db.Debug().Where("id = ?", IDdelete).Delete(&User{})

}
