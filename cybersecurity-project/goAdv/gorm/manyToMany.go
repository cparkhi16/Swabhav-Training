package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type User struct { //name user_id
	UserId  int `gorm:"primary_key;not null"`
	Name    string
	Courses []*Course `gorm:"many2many:user_courses;"`
}

type Course struct { //course_id name
	CourseId int `gorm:"primary_key;not null"`
	Name     string
	Users    []*User `gorm:"many2many:user_courses;"`
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

	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&Course{})
	/*

		course1 := Course{Name: "a", CourseId: 1}
		course2 := Course{Name: "b", CourseId: 2}

		user1 := User{Name: "c", UserId: 1, Courses: []*Course{&course1, &course2}}

		db.Debug().Create(&user1)*/

	var users []User
	err1 := db.Debug().Preload("Courses").Find(&users).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(users)

	var courses []Course
	db.Debug().Model(&User{UserId: 1}).Related(&courses, "Courses")
	fmt.Println(courses)
}
