package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Course struct {
	TestModel
	Name  string
	Users []*User `gorm:"many2many:person_courses;"`
}

func NewCourse(Name string) *Course {
	return &Course{Name: Name, TestModel: TestModel{CreatedAt: time.Now(), CreatedBy: "C", ID: uuid.NewV4()}}
}
