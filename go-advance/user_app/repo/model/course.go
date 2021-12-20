package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Course struct {
	TestModel
	Name  string
	Users []*User `gorm:"many2many:person_courses;"`
}

func NewCourse(Name string) *Course {
	return &Course{Name: Name, TestModel: TestModel{CreatedAtTime: time.Now(), CreatedBy: "C", ID: uuid.NewV4()}}
}
func (c *Course) BeforeCreate(scope *gorm.Scope) (err error) {
	c.CreatedAtTime = time.Now()
	c.CreatedBy = "Chinmay"
	c.ID = uuid.NewV4()
	scope.SetColumn("id", c.ID)
	scope.SetColumn("created_at_time", c.CreatedAtTime)
	scope.SetColumn("created_by", c.CreatedBy)
	fmt.Println("Setting Course uuid")
	return
}
