package model

import (
	"fmt"
	"time"
	"math/big"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type Publickey struct { //hobby_id name user_id
	Base
	N big.Int `gorm:"type:varbinary(500)"`
	E int 
	UserId uuid.UUID `gorm:"type:varchar(36)"`
}

func (h *Publickey) BeforeCreate(scope *gorm.Scope) (err error) {
	fmt.Println("here in publicKey")
	scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("CreatedBy", "yogesh")
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}
