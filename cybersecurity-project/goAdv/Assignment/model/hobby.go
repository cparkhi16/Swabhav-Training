package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type Hobby struct { //hobby_id name user_id
	CustomModel
	UserId uuid.UUID `gorm:"type:varchar(36);"`
	Name   string
}
