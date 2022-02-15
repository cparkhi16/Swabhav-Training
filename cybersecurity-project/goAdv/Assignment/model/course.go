package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Course struct { // course_id name user_id
	CustomModel
	Name  string
	Users []User `gorm:"many2many:user_courses;"`
}
