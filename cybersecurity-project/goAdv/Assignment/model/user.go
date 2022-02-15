package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type User struct { //user_id name address
	CustomModel
	Name    string
	Address string
	Courses []Course `gorm:"association_autoupdate:false;association_autocreate:false;many2many:user_courses;"`
	Hobbies []Hobby
}

func (u *User) BeforeCreate(scope *gorm.Scope) (err error) {
	fmt.Println("yes")
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
