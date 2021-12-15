package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	TestModel
	Name     string
	Address  string    `gorm:"column:ADDR"`
	Courses  []*Course `gorm:"many2many:person_courses;association_autoupdate:false;association_autocreate:false"`
	Hobbies  []Hobby
	Passport Passport
}

func NewUser(Name, Address string) *User {
	fmt.Println("Time created in NewUser-----------> ", time.Now())
	return &User{Name: Name, Address: Address, TestModel: TestModel{CreatedBy: "Chinmay", CreatedAt: time.Now()}}
}

func (u *User) AddHobbies(h Hobby) {
	u.Hobbies = append(u.Hobbies, h)
}

func (u *User) BeforeCreate() (err error) {
	u.ID = uuid.NewV4()
	if u.Name == "" {
		err = errors.New("can't save invalid data")
	}
	return
}

func (u *User) AfterCreate(scope *gorm.Scope) (err error) {
	fmt.Println("User created successfully")
	return
}

func (u *User) BeforeUpdate() (err error) {
	fmt.Println("-- Before update fired ---")
	return
}

func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	fmt.Println("-- After update fired ---")
	return
}

func (u *User) SetPassportForUser(p *Passport) {
	u.Passport = *p
}
