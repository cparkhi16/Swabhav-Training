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
	PassportID uint      `gorm:"unique"`
	ExpiryDate string
}

func NewPassport(PassportID uint) *Passport {
	return &Passport{PassportID: PassportID, TestModel: TestModel{CreatedBy: "J", CreatedAtTime: time.Now()}}
}

func (p *Passport) BeforeCreate(scope *gorm.Scope) (err error) {
	p.ID = uuid.NewV4()
	p.CreatedAtTime = time.Now()
	p.CreatedBy = "Chinmay"
	scope.SetColumn("id", p.ID)
	scope.SetColumn("created_at_time", p.CreatedAtTime)
	scope.SetColumn("created_by", p.CreatedBy)
	fmt.Println("Setting passport uuid")
	return
}
