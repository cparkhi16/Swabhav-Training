package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Hobby struct {
	TestModel
	UserID    uuid.UUID `gorm:"type:varchar(36)"`
	HobbyName string
}

func NewHobby(n string) *Hobby {
	fmt.Println("Created at time for new hobby for Rohit ----> ", time.Now())
	return &Hobby{HobbyName: n, TestModel: TestModel{ID: uuid.NewV4(), CreatedBy: "C", CreatedAt: time.Now()}}
}
func (h *Hobby) BeforeCreate(scope *gorm.Scope) (err error) {
	h.ID = uuid.NewV4()
	scope.SetColumn("id", h.ID)
	fmt.Println("Setting hobby uuid")
	return
}
