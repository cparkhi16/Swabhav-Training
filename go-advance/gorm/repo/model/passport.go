package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Passport struct {
	TestModel
	UserID     uuid.UUID `gorm:"type:varchar(36)"`
	PassportID uint
}

func NewPassport(PassportID uint) *Passport {
	return &Passport{PassportID: PassportID, TestModel: TestModel{CreatedBy: "J", CreatedAt: time.Now(), ID: uuid.NewV4()}}
}

func (p *Passport) BeforeCreate(scope *gorm.Scope) (err error) {
	p.ID = uuid.NewV4()
	scope.SetColumn("id", p.ID)
	fmt.Println("Setting passport uuid")
	return
}
