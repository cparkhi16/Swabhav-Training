package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type Hobby struct { //hobby_id name user_id
	Base
	Name   string    `json:",omitempty"`
	UserId uuid.UUID `gorm:"type:varchar(36);" json:",omitempty"`
}

func (h *Hobby) BeforeCreate(scope *gorm.Scope) (err error) {
	fmt.Println("here in hobby")
	scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("CreatedBy", "yogesh")
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}
