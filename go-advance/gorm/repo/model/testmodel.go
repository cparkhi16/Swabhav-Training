package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TestModel struct {
	CreatedBy string
	CreatedAt time.Time //This is same as gorm.model so while updating in db it will take it's own value
	DeletedAt *time.Time
	ID        uuid.UUID `gorm:"type:varchar(36);primary_key"`
}
