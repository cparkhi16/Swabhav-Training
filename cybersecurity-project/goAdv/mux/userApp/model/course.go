package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type Course struct {
	Base
	Name  string `gorm:"unique;not null" json:",omitempty"`
	Users []User `gorm:"many2many:user_courses;" json:",omitempty"`
	Price int    `gorm:"type:int;" json:"Price,omitempty"`
}

func (u *Course) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("CreatedBy", "yogesh")
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}
