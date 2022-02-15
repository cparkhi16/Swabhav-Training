package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:varchar(36);primary_key" json:",omitempty"`
	CreatedBy string     `json:",omitempty"`
	CreatedAt time.Time  `json:",omitempty"`
	DeletedAt *time.Time `json:",omitempty"`
}
