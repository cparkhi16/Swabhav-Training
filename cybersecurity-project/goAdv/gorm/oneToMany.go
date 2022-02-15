package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Gamer struct { //name usrid
	Name    string
	UserId  int     `gorm:"primary_key;not null"`
	Hobbies []Hobby `gorm:"foreignkey:GamerId"`
}

type Hobby struct { //hobbyid (pk) gamerid name
	GamerId int
	HobbyId int `gorm:"primary_key;not null"`
	Name    string
}

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

	/*
		db.AutoMigrate(&Gamer{}, &Hobby{})
		err2 := db.Debug().Model(&Hobby{}).AddForeignKey("gamer_id", "gamers(user_id)", "CASCADE", "CASCADE").Error
		if err2 != nil {
			fmt.Println(err2)
		}*/

	h := Hobby{HobbyId: 34, Name: "drawing"}
	h2 := Hobby{HobbyId: 13, Name: "ster"}
	var list []Hobby
	list = append(list, h)
	list = append(list, h2)

	g := Gamer{UserId: 3}
	//db.Debug().Create(&g)
	fmt.Println(g)

	db.Debug().Delete(&g)

}
