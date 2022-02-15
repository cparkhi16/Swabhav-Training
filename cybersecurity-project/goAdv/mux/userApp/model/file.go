package model

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type File struct {
	Base
	FileName  string `gorm:"unique;not null" json:",omitempty"`
	LevelBIBA int    `gorm:"type:int;" json:"LevelBIBA,omitempty"`
	LevelBell int    `gorm:"type:int;" json:"LevelBell,omitempty"`
}

// func NewFile(name, bell, biba string) *File {
//     return &File{FileName: name, LevelBIBA: biba, LevelBell: bell}
// }

func (u *File) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("CreatedBy", "yogesh")
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}
