package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CustomModel struct {
	ID        uuid.UUID `gorm:"type:varchar(36);primary_key"`
	CreatedBy string
	CreatedAt time.Time
	DeletedAt *time.Time
}
