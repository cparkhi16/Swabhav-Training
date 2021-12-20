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
	return &Hobby{HobbyName: n, TestModel: TestModel{ID: uuid.NewV4(), CreatedBy: "C", CreatedAtTime: time.Now()}}
}
func (h *Hobby) BeforeCreate(scope *gorm.Scope) (err error) {
	h.CreatedAtTime = time.Now()
	h.CreatedBy = "Chinmay"
	h.ID = uuid.NewV4()
	scope.SetColumn("id", h.ID)
	scope.SetColumn("created_at_time", h.CreatedAtTime)
	scope.SetColumn("created_by", h.CreatedBy)
	fmt.Println("Setting hobby uuid")
	return
}
