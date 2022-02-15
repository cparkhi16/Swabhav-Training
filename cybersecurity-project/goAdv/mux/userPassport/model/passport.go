package model

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type Passport struct {
	ID      uuid.UUID `gorm:"primary_key;type:varchar(36)"`
	Country string
	UserId  uuid.UUID `gorm:"type:varchar(36)"`
}

func (u *Passport) BeforeCreate(scope *gorm.Scope) (err error) {
	//fmt.Println("yes passport")
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
