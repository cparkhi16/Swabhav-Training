package model

import (
	"net/mail"
	"time"
	//"crypto/rsa"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Base
	FirstName string   `json:",omitempty"`
	LastName  string   `json:",omitempty"`
	Email     string   `gorm:"unique;not null"`
	Password  string   `gorm:"not null"`
	Passport  Passport `gorm:"foreignkey:UserId;" json:",omitempty"`
	Courses   []Course `gorm:"many2many:user_courses;constraint:OnDelete:CASCADE",json:"Courses,omitempty"`
	Hobbies   []Hobby  `json:"Hobbies,omitempty"`
	LevelBIBA int      `gorm:"type:int;" json:"LevelBIBA,omitempty"`
	LevelBell int      `gorm:"type:int;" json:"LevelBell,omitempty"`
	SecretAnswer string `gorm:"type:varbinary(5000);" json:",omitempty"`
	//Publickey  Publickey `gorm:"foreignkey:UserId;" json:",omitempty"`
	PrivateKey string `gorm:"type:varchar(5000)"json:",omitempty"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("CreatedBy", "yogesh")
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

func (u *User) Validate() bool {
	_, err := mail.ParseAddress(u.Email)
	if err == nil {
		return true
	}
	return false
}
