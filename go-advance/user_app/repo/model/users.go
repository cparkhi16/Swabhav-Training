package model

import (
	lr "app/logger"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Login struct {
	Email    string
	Password string
}
type User struct {
	TestModel
	FirstName string
	LastName  string
	Address   string `gorm:"column:ADDR"`
	Email     string `gorm:"unique"`
	Password  string
	Courses   []*Course `gorm:"many2many:person_courses;association_autoupdate:false;association_autocreate:false"`
	Hobbies   []Hobby
	Passport  Passport
}

func NewUser(FirstName, LastName, Address string) *User {
	fmt.Println("Created at time ", time.Now())
	return &User{FirstName: FirstName, LastName: LastName, Address: Address, TestModel: TestModel{CreatedBy: "Chinmay", CreatedAtTime: time.Now()}}
}

func (u *User) AddHobbies(h Hobby) {
	u.Hobbies = append(u.Hobbies, h)
}

func (u *User) BeforeCreate(scope *gorm.Scope) (err error) {
	u.ID = uuid.NewV4()
	u.CreatedAtTime = time.Now()
	u.CreatedBy = "Chinmay"
	//fmt.Println("Assigning ID TO user", u.ID)
	//scope.SetColumn("id", uuid.NewV4())
	if (u.FirstName == "") || (u.LastName == "") || (u.Email == "") || (u.Password == "") {
		logger := lr.GetLogger()
		logger.Debug().Msg("Empty field found in user creation")
		err = errors.New("can't save invalid data")
	}
	return
}

func (u *User) AfterCreate(scope *gorm.Scope) (err error) {
	logger := lr.GetLogger()
	logger.Debug().Msg("User created successfully")
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
