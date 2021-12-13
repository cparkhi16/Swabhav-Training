package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TestModel struct {
	CreatedBy string
	CreatedAt time.Time
	DeletedAt *time.Time
	ID        uuid.UUID `gorm:"type:varchar(36);primary_key"`
}
