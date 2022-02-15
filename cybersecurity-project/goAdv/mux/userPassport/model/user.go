package model

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primary_key;type:varchar(36)"`
	Name     string
	Passport Passport `gorm:"foreignkey:UserId"`
}

/*

type UserDTO struct{
	ID       uuid.UUID `gorm:"primary_key;type:varchar(36)"`
	Name     string
	Passport Passport `gorm:"foreignkey:UserId"`
	Course   []Course `gorm:"foreignkey:UserId"`
}*/

func (u *User) BeforeCreate(scope *gorm.Scope) (err error) {
	fmt.Println("yes")
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
