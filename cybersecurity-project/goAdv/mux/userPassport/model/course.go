package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type Course struct {
	ID     uuid.UUID `gorm:"primary_key;type:varchar(36)"`
	Name   string
	UserId uuid.UUID
}

func (u *Course) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
