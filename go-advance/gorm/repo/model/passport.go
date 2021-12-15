package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Passport struct {
	TestModel
	UserID     uuid.UUID `gorm:"type:varchar(36)"`
	PassportID uint
}

func NewPassport(PassportID uint) *Passport {
	return &Passport{PassportID: PassportID, TestModel: TestModel{CreatedBy: "Chinmay", CreatedAt: time.Now(), ID: uuid.NewV4()}}
}

/*func (p *Passport) BeforeCreate() (err error) {
	p.ID = uuid.NewV4()
	fmt.Println("Setting passport uuid")
	return
}*/
