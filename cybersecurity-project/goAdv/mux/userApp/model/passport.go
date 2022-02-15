package model

import (
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type Passport struct {
	Base
	PassportNo uuid.UUID `gorm:"type:varchar(36);" json:",omitempty"`
	ExpiryDate string    `gorm:"type:datetime;" json:",omitempty"`
	Country    string    `json:",omitempty"`
	UserId     uuid.UUID `gorm:"type:varchar(36)" json:",omitempty"`
}

func (u *Passport) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("PassportNo", uuid.NewV4())
	scope.SetColumn("CreatedBy", "yogesh")
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}
